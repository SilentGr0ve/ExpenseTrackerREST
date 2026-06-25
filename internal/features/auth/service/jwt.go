package auth_service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) generateAccessToken(ID uuid.UUID, email string) (string, error) {
	claims := jwt.MapClaims{
		"subject":   ID,
		"email":     email,
		"expiresAt": time.Now().Add(s.jwtExpiry * time.Minute).Unix(),
		"issuedAt":  time.Now().Unix(),
		"issuer":    "expense-tracker",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func checkPasswordHash(password string, hash []byte) error {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
