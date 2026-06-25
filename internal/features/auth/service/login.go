package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
)

func (s *AuthService) Login(ctx context.Context, req domain.LoginRequest) (string, error) {

	user, err := s.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return "", core_errors.ErrUnauthorized
		}
		return "", fmt.Errorf("get user from repository: %w", err)
	}

	if err := checkPasswordHash(req.Password, []byte(user.PasswordHash)); err != nil {
		return "", fmt.Errorf("compare hash and password: %w", err)
	}

	accessToken, err := s.generateAccessToken(user.ID, user.Email)
	if err != nil {
		return "", fmt.Errorf("generating token: %w", err)
	}

	return accessToken, nil
}
