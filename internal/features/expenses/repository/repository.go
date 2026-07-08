package expenses_repository

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/repository"
)

type ExpensesRepository struct {
	pool      repository.Pool
	opTimeout time.Duration
}

func NewExpensesRepository(pool repository.Pool, opTimeout time.Duration) *ExpensesRepository {
	return &ExpensesRepository{
		pool:      pool,
		opTimeout: opTimeout,
	}
}
