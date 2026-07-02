package categories_transport

import (
	"context"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
	"github.com/google/uuid"
)

type CategoriesHTTPHandler struct {
	categoriesService CategoriesService
}

type CategoriesService interface {
	CreateCategory(ctx context.Context, userID uuid.UUID, name string) (domain.Category, error)
	GetCategories(ctx context.Context, userID uuid.UUID) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
	UpdateCategory(ctx context.Context, id uuid.UUID, userID uuid.UUID, patch domain.CategoryPatch) (domain.Category, error)
}

func NewCategoriesHTTPHandler(categoriesService CategoriesService) *CategoriesHTTPHandler {
	return &CategoriesHTTPHandler{
		categoriesService: categoriesService,
	}
}

func (h *CategoriesHTTPHandler) ProtectedRoutes() []httpserver.Route {
	return []httpserver.Route{
		{Method: http.MethodPost, Path: "/categories", Handler: h.CreateCategory},
		{Method: http.MethodGet, Path: "/categories", Handler: h.GetCategories},
		{Method: http.MethodDelete, Path: "/categories/{id}", Handler: h.DeleteCategory},
		{Method: http.MethodPatch, Path: "/categories/{id}", Handler: h.UpdateCategory},
	}
}
