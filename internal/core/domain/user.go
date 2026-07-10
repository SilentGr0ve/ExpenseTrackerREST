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

type UserPatch struct {
	FullName *string
	Email    *string
	Password *string
}
