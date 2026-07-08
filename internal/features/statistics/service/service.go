package statistics_service

import (
	"context"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetStatistics(ctx context.Context, userID uuid.UUID, dateFrom, dateTo *time.Time) (domain.Statistics, error)
}

func NewStatisticsService(statisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: statisticsRepository,
	}
}
