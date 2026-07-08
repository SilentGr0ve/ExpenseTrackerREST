package statistics_transport

import (
	"context"
	"net/http"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
	"github.com/google/uuid"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(ctx context.Context, userID uuid.UUID, dateFrom, dateTo *time.Time) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) ProtectedRoutes() []httpserver.Route {
	return []httpserver.Route{
		{Method: http.MethodGet, Path: "/statistics", Handler: h.GetStatistics},
	}
}
