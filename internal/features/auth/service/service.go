package auth_service

import (
	"context"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
)

type AuthService struct {
	authRepository AuthRepositoryInterface
}

type AuthRepositoryInterface interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}
