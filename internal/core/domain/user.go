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

type SensitiveUserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Version   int        `json:"version"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
