package categories_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
)

func (h *CategoriesHTTPHandler) DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	id := r.PathValue("id")
	categoryID, err := uuid.Parse(id)
	if err != nil {
		rh.ErrorResponse(core_errors.ErrInvalidArgument, "failed to parse id")
		return
	}

	if err := h.categoriesService.DeleteCategory(ctx, categoryID, userID); err != nil {
		rh.ErrorResponse(err, "failed to delete category")
		return
	}

	rh.NoContent()
}
