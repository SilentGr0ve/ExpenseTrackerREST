package auth_transport

import (
	"context"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
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
