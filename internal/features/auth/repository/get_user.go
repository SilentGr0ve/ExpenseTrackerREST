package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	SELECT id, version, full_name, email, created_at, updated_at FROM tracker.users WHERE email = $1; 
	`

	row := r.pool.QueryRow(ctx, query, email)

	var user domain.User

	if err := row.Scan(
		&user.ID,
		&user.Version,
		&user.FullName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with email='%s': %w", email, core_errors.ErrNotFound)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return user, nil
}
