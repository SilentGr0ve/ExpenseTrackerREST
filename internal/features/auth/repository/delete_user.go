package auth_repository

import (
	"context"
	"fmt"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (r *AuthRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	DELETE FROM tracker.users WHERE id = $1;
	`

	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
