package auth_transport

import (
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	var loginRequest domain.LoginRequest

	if err := request.DecodeAndValidateRequest(r, &loginRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	token, err := h.authService.Login(ctx, loginRequest)
	if err != nil {
		rh.ErrorResponse(err, "failed to authorize")
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken: token,
	}

	rh.JSONResponse(http.StatusOK, loginResponse)
}
