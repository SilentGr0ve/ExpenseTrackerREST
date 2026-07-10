package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, login domain.Login) (string, error) {

	user, err := s.authRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return "", core_errors.ErrUnauthorized
		}
		return "", fmt.Errorf("get user from repository: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password)); err != nil {
		return "", core_errors.ErrUnauthorized
	}

	accessToken, err := s.generateAccessToken(user.ID, user.Email)
	if err != nil {
		return "", fmt.Errorf("generating token: %w", err)
	}

	return accessToken, nil
}
