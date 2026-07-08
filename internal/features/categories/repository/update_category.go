package categories_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *CategoriesRepository) UpdateCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID, patch domain.CategoryPatch) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	UPDATE tracker.categories
	SET name=$1,version=version+1, updated_at=NOW()
	WHERE id=$2 AND user_id=$3
	RETURNING id, version,user_id, name, created_at, updated_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		*patch.Name,
		id,
		userID,
	)

	var updated domain.Category

	if err := row.Scan(
		&updated.ID,
		&updated.Version,
		&updated.UserID,
		&updated.Name,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Category{}, core_errors.ErrConflict
		}
		return domain.Category{}, fmt.Errorf("scan updated category: %w", err)
	}

	return updated, nil
}
