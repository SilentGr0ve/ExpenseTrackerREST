package categories_service

import (
	"context"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

func (s *CategoriesService) GetCategories(ctx context.Context, userID uuid.UUID) ([]domain.Category, error) {
	categories, err := s.categoriesRepository.GetCategories(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get categories from repository: %w", err)
	}
	return categories, nil
}
