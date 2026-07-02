package categories_transport

import (
	"errors"
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

func (h *CategoriesHTTPHandler) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var categoryRequest CreateCategoryRequest

	if err := request.DecodeAndValidateRequest(r, &categoryRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	category, err := h.categoriesService.CreateCategory(ctx, userID, categoryRequest.Name)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			rh.ErrorResponse(err, "category already exists")
			return
		}
		rh.ErrorResponse(err, "failed to create category")
		return
	}

	categoryResponse := CategoryResponse{
		ID:        category.ID,
		Version:   category.Version,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	rh.JSONResponse(
		http.StatusCreated,
		categoryResponse,
	)
}
