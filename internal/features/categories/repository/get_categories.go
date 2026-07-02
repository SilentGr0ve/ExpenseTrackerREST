package categories_repository

import (
	"context"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

func (r *CategoriesRepository) GetCategories(ctx context.Context, userID uuid.UUID) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	SELECT id, version,user_id,name,created_at, updated_at
	FROM tracker.categories 
	WHERE user_id=$1
	ORDER BY created_at DESC;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("get categories: %w", err)
	}

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(
			&category.ID,
			&category.Version,
			&category.UserID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan categories: %w", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return categories, nil
}
