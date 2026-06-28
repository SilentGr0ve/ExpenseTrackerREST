package domain

import "time"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,min=8,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=8,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type RegisterResponse struct {
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
