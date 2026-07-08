package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (s *StatisticsService) GetStatistics(ctx context.Context, userID uuid.UUID, dateFrom, dateTo *time.Time) (domain.Statistics, error) {
	if dateFrom != nil && dateTo != nil {
		if dateTo.Before(*dateFrom) {
			return domain.Statistics{}, fmt.Errorf("date_to can't be before date_from: %w", core_errors.ErrInvalidArgument)
		}
	}

	statistics, err := s.statisticsRepository.GetStatistics(ctx, userID, dateFrom, dateTo)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get statistics from repository: %w", err)
	}

	return statistics, nil
}
