package auth_transport

import (
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

type updateUserRequest struct {
	FullName *string `json:"full_name" validate:"omitempty,min=3,max=100"`
	Email    *string `json:"email"     validate:"omitempty,email"`
	Password *string `json:"password"  validate:"omitempty,min=6,max=255"`
}

func (h *AuthHTTPHandler) UpdateMe(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var patchRequest updateUserRequest
	if err := request.DecodeAndValidateRequest(r, &patchRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	patch := domain.UserPatch{
		FullName: patchRequest.FullName,
		Email:    patchRequest.Email,
		Password: patchRequest.Password,
	}

	user, err := h.authService.UpdateUser(ctx, userID, patch)
	if err != nil {
		rh.ErrorResponse(err, "failed to update user")
		return
	}

	patchedUser := domain.UserResponse{
		ID:        user.ID,
		Version:   user.Version,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	rh.JSONResponse(
		http.StatusOK,
		patchedUser,
	)
}
