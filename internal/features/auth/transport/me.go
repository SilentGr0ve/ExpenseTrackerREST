package auth_transport

import (
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

func (h *AuthHTTPHandler) GetMe(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		rh.ErrorResponse(err, "failed to get user")
		return
	}

	getUser := domain.UserResponse{
		ID:        user.ID,
		Version:   user.Version,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	rh.JSONResponse(http.StatusOK, getUser)
}
