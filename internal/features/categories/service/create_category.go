package categories_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (s *CategoriesService) CreateCategory(ctx context.Context, userID uuid.UUID, name string) (domain.Category, error) {
	category := domain.Category{
		ID:        uuid.New(),
		Version:   1,
		UserID:    userID,
		Name:      name,
		CreatedAt: time.Now(),
	}

	category, err := s.categoriesRepository.CreateCategory(ctx, category)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			return domain.Category{}, fmt.Errorf("category already exists: %w", err)
		}
		return domain.Category{}, fmt.Errorf("create category in repository: %w", err)
	}

	return category, nil
}
