package categories_repository

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/repository"
)

type CategoriesRepository struct {
	pool      repository.Pool
	opTimeout time.Duration
}

func NewCategoriesRepository(pool repository.Pool, opTimeout time.Duration) *CategoriesRepository {
	return &CategoriesRepository{
		pool:      pool,
		opTimeout: opTimeout,
	}
}
