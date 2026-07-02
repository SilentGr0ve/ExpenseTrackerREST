package categories_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (s *CategoriesService) UpdateCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID, patch domain.CategoryPatch) (domain.Category, error) {
	if patch.Name == nil {
		return domain.Category{}, core_errors.ErrInvalidArgument
	}

	updatedCategory, err := s.categoriesRepository.UpdateCategory(ctx, id, userID, patch)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			return domain.Category{}, fmt.Errorf("category modified earlier: %w", err)
		}
		return domain.Category{}, fmt.Errorf("update category in repository: %w", err)
	}

	return updatedCategory, nil
}
