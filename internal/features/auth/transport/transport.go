package auth_transport

import (
	"context"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
	"github.com/google/uuid"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Register(ctx context.Context, req domain.RegisterRequest) (domain.User, error)
	Login(ctx context.Context, req domain.LoginRequest) (string, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UpdateUser(ctx context.Context, id uuid.UUID, patch domain.UserPatch) (domain.User, error)
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
	return []httpserver.Route{
		{Method: http.MethodGet, Path: "/users/me", Handler: h.GetMe},
		{Method: http.MethodDelete, Path: "/users/me", Handler: h.DeleteMe},
		{Method: http.MethodPatch, Path: "/users/me", Handler: h.UpdateMe},
	}
}
