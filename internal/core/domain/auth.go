package domain

import "time"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,min=8,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
