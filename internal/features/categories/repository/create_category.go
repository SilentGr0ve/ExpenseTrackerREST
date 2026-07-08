package categories_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CategoriesRepository) CreateCategory(ctx context.Context, category domain.Category) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	INSERT INTO tracker.categories (id, version, user_id, name, created_at)
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id, version, user_id, name, created_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		category.ID,
		category.Version,
		category.UserID,
		category.Name,
		category.CreatedAt,
	)

	var newCategory domain.Category

	if err := row.Scan(
		&newCategory.ID,
		&newCategory.Version,
		&newCategory.UserID,
		&newCategory.Name,
		&newCategory.CreatedAt,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.Category{}, core_errors.ErrConflict
		}
		return domain.Category{}, fmt.Errorf("scan category model: %w", err)
	}

	return newCategory, nil
}
