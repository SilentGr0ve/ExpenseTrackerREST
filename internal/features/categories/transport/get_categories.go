package categories_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

// GetCategories godoc
// @Summary Get user categories
// @Description Get authorized user's categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {array} CategoryResponse "User categories returned"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /categories [get]
func (h *CategoriesHTTPHandler) GetCategories(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	categories, err := h.categoriesService.GetCategories(ctx, userID)
	if err != nil {
		rh.ErrorResponse(err, "failed to get categories")
		return
	}

	categoriesResponse := ToDTOFromDomain(categories)

	rh.JSONResponse(
		http.StatusOK,
		categoriesResponse,
	)

}
