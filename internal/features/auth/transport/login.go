package auth_transport

import (
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

// LoginUser godoc
// @Summary Login user
// @Description Authenticate with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "Access token returned"
// @Failure 400 {object} response.ErrorResponse "Invalid request body"
// @Failure 401 {object} response.ErrorResponse "Invalid credentials"
// @Failure 404 {object} response.ErrorResponse "User not found"
// @Router /auth/login [post]
func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	var loginRequest LoginRequest

	if err := request.DecodeAndValidateRequest(r, &loginRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	login := domain.Login{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	token, err := h.authService.Login(ctx, login)
	if err != nil {
		rh.ErrorResponse(err, "failed to authorize")
		return
	}

	loginResponse := LoginResponse{
		AccessToken: token,
	}

	rh.JSONResponse(http.StatusOK, loginResponse)
}
