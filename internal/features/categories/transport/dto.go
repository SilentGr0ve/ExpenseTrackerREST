package categories_transport

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100" example:"Fastfood"`
}

type CategoryResponse struct {
	ID        uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Version   int        `json:"version" example:"1"`
	Name      string     `json:"name" example:"Fastfood"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CategoryPatchRequest struct {
	Name *string `json:"name" validate:"omitempty,min=3,max=100" example:"John Doe"`
}

func ToDTOFromDomain(domainCategories []domain.Category) []CategoryResponse {
	categoriesResponse := make([]CategoryResponse, 0, len(domainCategories))
	for _, domainCategory := range domainCategories {
		categoriesResponse = append(
			categoriesResponse,
			CategoryResponse{
				ID:        domainCategory.ID,
				Version:   domainCategory.Version,
				Name:      domainCategory.Name,
				CreatedAt: domainCategory.CreatedAt,
				UpdatedAt: domainCategory.UpdatedAt,
			})
	}

	return categoriesResponse
}
