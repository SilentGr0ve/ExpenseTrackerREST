package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID
	Version int

	FullName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func NewUser(id uuid.UUID, version int, fullName, email, passwordHash string, createdAt time.Time) User {
	return User{
		ID:           id,
		Version:      version,
		FullName:     fullName,
		Email:        email,
		PasswordHash: passwordHash,
	}
}

func CreateDomainUser(fullName, email, passwordHash string) User {
	var (
		id        = uuid.New()
		version   = 1
		createdAt = time.Now()
	)

	return NewUser(
		id,
		version,
		fullName,
		email,
		passwordHash,
		createdAt,
	)
}
