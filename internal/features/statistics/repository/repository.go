package statistics_repository

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/repository"
)

type StatisticsRepository struct {
	pool      repository.Pool
	opTimeout time.Duration
}

func NewStatisticsRepository(pool repository.Pool, opTimeout time.Duration) *StatisticsRepository {
	return &StatisticsRepository{
		pool:      pool,
		opTimeout: opTimeout,
	}
}
