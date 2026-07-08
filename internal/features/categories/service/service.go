package categories_service

import (
	"context"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

type CategoriesService struct {
	categoriesRepository CategoriesRepository
}

type CategoriesRepository interface {
	CreateCategory(ctx context.Context, req domain.Category) (domain.Category, error)
	GetCategories(ctx context.Context, userID uuid.UUID) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
	UpdateCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID, patch domain.CategoryPatch) (domain.Category, error)
}

func NewCategoryService(categoriesRepository CategoriesRepository) *CategoriesService {
	return &CategoriesService{
		categoriesRepository: categoriesRepository,
	}
}
