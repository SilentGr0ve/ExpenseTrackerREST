package auth_transport

import (
	"errors"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with email, password and full name.
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration Payload"
// @Success 201 {object} RegisterResponse "User created successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid request body"
// @Failure 409 {object} response.ErrorResponse "Email already exists"
// @Router /auth/register [post]
func (h *AuthHTTPHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	var registerRequest RegisterRequest

	if err := request.DecodeAndValidateRequest(r, &registerRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	register := domain.Register{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		FullName: registerRequest.FullName,
	}

	user, err := h.authService.Register(ctx, register)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			rh.ErrorResponse(err, "user with this email already exists")
			return
		}
		rh.ErrorResponse(err, "failed to create user")
		return
	}

	registerResponse := RegisterResponse{
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}

	rh.JSONResponse(http.StatusCreated, registerResponse)
}
