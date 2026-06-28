package auth_transport

import (
	"context"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Register(ctx context.Context, req domain.RegisterRequest) (domain.User, error)
	Login(ctx context.Context, req domain.LoginRequest) (string, error)
}

func NewAuthHTTPHandler(authService AuthService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
	}
}

func (h *AuthHTTPHandler) Routes() []httpserver.Route {
	return []httpserver.Route{
		{Method: http.MethodPost, Path: "/auth/register", Handler: h.Register},
		{Method: http.MethodPost, Path: "/auth/login", Handler: h.Login},
	}
}

func (h *AuthHTTPHandler) ProtectedRoutes() []httpserver.Route {
	return []httpserver.Route{}
}
