package auth_repository

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/repository"
)

type AuthRepository struct {
	pool      repository.Pool
	opTimeout time.Duration
}

func NewAuthRepository(pool repository.Pool, opTimeout time.Duration) *AuthRepository {
	return &AuthRepository{
		pool:      pool,
		opTimeout: opTimeout,
	}
}
