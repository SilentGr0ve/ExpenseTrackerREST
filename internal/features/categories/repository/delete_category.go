package categories_repository

import (
	"context"
	"fmt"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (r *CategoriesRepository) DeleteCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	DELETE FROM tracker.categories WHERE id=$1 AND user_id=$2;
	`
	tag, err := r.pool.Exec(
		ctx,
		query,
		id,
		userID,
	)
	if err != nil {
		return fmt.Errorf("delete category: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
