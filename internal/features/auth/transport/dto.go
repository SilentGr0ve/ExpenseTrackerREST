package auth_transport

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,min=8,max=255" example:"johndoel@mail.com"`
	Password string `json:"password" validate:"required,min=6,max=255" example:"johndoepassword"`
	FullName string `json:"full_name" validate:"required,min=3,max=100" example:"John Doe"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=8,max=255" example:"johndoel@mail.com"`
	Password string `json:"password" validate:"required,min=6,max=255" example:"johndoepassword"`
}

type RegisterResponse struct {
	Email     string    `json:"email" example:"johndoel@mail.com"`
	FullName  string    `json:"full_name" example:"John Doe"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

type UserResponse struct {
	ID        uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Version   int        `json:"version" example:"1"`
	FullName  string     `json:"full_name" example:"John Doe"`
	Email     string     `json:"email" example:"johndoe@mail.com"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UpdateUserRequest struct {
	FullName *string `json:"full_name" validate:"omitempty,min=3,max=100" example:"John Doe"`
	Email    *string `json:"email"     validate:"omitempty,email" example:"johndoe@mail.com"`
	Password *string `json:"password"  validate:"omitempty,min=6,max=255" example:"johndoepassword"`
}
