package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	auth_transport "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/auth/transport"
	"github.com/google/uuid"
)

func (s *AuthService) UpdateUser(ctx context.Context, id uuid.UUID, patch auth_transport.UserPatch) (domain.User, error) {
	user, err := s.authRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if patch.FullName != nil {
		user.FullName = *patch.FullName
	}
	if patch.Email != nil {
		user.Email = *patch.Email
	}
	if patch.Password != nil {
		hash, err := hashPassword(*patch.Password)
		if err != nil {
			return domain.User{}, fmt.Errorf("hash password: %w", err)
		}
		user.PasswordHash = hash
	}

	updatedUser, err := s.authRepository.UpdateUser(ctx, user)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			return domain.User{}, fmt.Errorf("user modified earlier: %w", err)
		}
		return domain.User{}, fmt.Errorf("update user in repository: %w", err)
	}

	return updatedUser, nil
}
