package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
)

func (s *AuthService) Register(ctx context.Context, req domain.RegisterRequest) (domain.User, error) {
	hash, err := hashPassword(req.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hashing password: %w", err)
	}

	userDomain := domain.CreateDomainUser(
		req.FullName,
		req.Email,
		hash,
	)

	user, err := s.authRepository.CreateUser(ctx, userDomain)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			return domain.User{}, fmt.Errorf("user already exists: %w", err)
		}
		return domain.User{}, fmt.Errorf("create user in repository: %w", err)
	}

	return user, nil
}
