package auth_service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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
