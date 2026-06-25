package auth_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, req domain.RegisterRequest) (domain.User, error) {
	hash, err := hashPassword(req.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hashing password: %w", err)
	}

	userDomain := domain.User{
		ID:           uuid.New(),
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: hash,
		Version:      1,
	}

	user, err := s.authRepository.CreateUser(ctx, userDomain)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			return domain.User{}, fmt.Errorf("user already exists: %w", err)
		}
		return domain.User{}, fmt.Errorf("create user in repository: %w", err)
	}

	return user, nil
}

func (s *AuthService) generateAccessToken(ID uuid.UUID, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   ID,
		"email": email,
		"exp":   time.Now().Add(s.jwtExpiry).Unix(),
		"iat":   time.Now().Unix(),
		"iss":   "expense-tracker",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
