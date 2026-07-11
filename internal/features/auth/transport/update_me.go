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

// UpdateMeUser godoc
// @Summary Update current user
// @Description Update the authenticated user's profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body UpdateUserRequest true "Partial user fields to update"
// @Success 200 {object} UserResponse "Updated user returned"
// @Failure 400 {object} response.ErrorResponse "Invalid request body"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /users/me [patch]
func (h *AuthHTTPHandler) UpdateMe(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var patchRequest UpdateUserRequest
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

	patchedUser := UserResponse{
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
