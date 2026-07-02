package categories_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *CategoriesService) DeleteCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	err := s.categoriesRepository.DeleteCategory(ctx, id, userID)
	if err != nil {
		return fmt.Errorf("delete user from repository: %w", err)
	}
	return nil
}
