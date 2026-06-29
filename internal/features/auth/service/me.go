package auth_service

import (
	"context"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

func (s *AuthService) GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.authRepository.GetUserByID(ctx, id)
}
