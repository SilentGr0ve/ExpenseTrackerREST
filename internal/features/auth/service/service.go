package auth_service

import (
	"context"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

type AuthService struct {
	authRepository AuthRepository
	jwtSecret      string
	jwtExpiry      time.Duration
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
}

func NewAuthService(authRepository AuthRepository, jwtSecret string, jwtExpiry time.Duration) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		jwtSecret:      jwtSecret,
		jwtExpiry:      jwtExpiry,
	}
}
