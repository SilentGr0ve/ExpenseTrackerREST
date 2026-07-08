package auth_service

import (
	"context"

	"github.com/google/uuid"
)

func (s *AuthService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.authRepository.DeleteUser(ctx, id)
}
