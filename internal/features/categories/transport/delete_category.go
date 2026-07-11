package categories_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
)

// DeleteCategory godoc
// @Summary Delete current category
// @Description Delete the authenticated user's category
// @Tags categories
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 204 "Category deleted successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid category id"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Category not found"
// @Failure 409 {object} response.ErrorResponse "The category is used by expenses"
// @Router /categories/{id} [delete]
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
