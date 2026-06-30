package auth_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *AuthRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	UPDATE tracker.users 
	SET full_name=$1, email=$2, password_hash=$3, updated_at=NOW(), version=version+1
	WHERE id=$4 AND version=$5
	RETURNING 	id, version, full_name, email, password_hash, created_at, updated_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.ID,
		user.Version,
	)

	var updated domain.User
	if err := row.Scan(
		&updated.ID,
		&updated.Version,
		&updated.FullName,
		&updated.Email,
		&updated.PasswordHash,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, core_errors.ErrConflict
		}
		return domain.User{}, fmt.Errorf("scan updated user: %w", err)
	}

	return updated, nil
}
