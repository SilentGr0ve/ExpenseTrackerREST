package auth_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

// DeleteMeUser godoc
// @Summary Delete current user
// @Description Delete the authenticated user's account
// @Tags users
// @Security BearerAuth
// @Success 204 "User deleted"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /users/me [delete]
func (h *AuthHTTPHandler) DeleteMe(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	if err := h.authService.DeleteUser(ctx, userID); err != nil {
		rh.ErrorResponse(err, "failed to delete user")
		return
	}

	rh.NoContent()
}
