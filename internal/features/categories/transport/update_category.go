package categories_transport

import (
	"errors"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
)

func (h *CategoriesHTTPHandler) UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		rh.ErrorResponse(core_errors.ErrInvalidArgument, "failed to parse id")
		return
	}

	var patchRequest CategoryPatchRequest
	if err := request.DecodeAndValidateRequest(r, &patchRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	patch := domain.CategoryPatch{
		Name: patchRequest.Name,
	}

	category, err := h.categoriesService.UpdateCategory(ctx, id, userID, patch)
	if err != nil {
		if errors.Is(err, core_errors.ErrInvalidArgument) {
			rh.ErrorResponse(err, "nothing to update")
			return
		}
		rh.ErrorResponse(err, "failed to patch category")
		return
	}

	patchedCategory := CategoryResponse{
		ID:        category.ID,
		Version:   category.Version,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	rh.JSONResponse(http.StatusOK, patchedCategory)
}
