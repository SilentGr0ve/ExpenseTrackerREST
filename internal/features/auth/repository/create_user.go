package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AuthRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	INSERT INTO tracker.users (id, version, full_name, email, password_hash, created_at, updated_at)  
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, version, full_name, email, password_hash, created_at, updated_at; 
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.ID,
		user.Version,
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)

	var newUser domain.User

	if err := row.Scan(&newUser.ID, &newUser.Version, &newUser.FullName, &newUser.Email, &newUser.PasswordHash, &newUser.CreatedAt, &newUser.UpdatedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.User{}, core_errors.ErrConflict
		}
		return domain.User{}, fmt.Errorf("scan user model: %w", err)
	}

	return newUser, nil
}
