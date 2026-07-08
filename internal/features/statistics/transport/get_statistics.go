package statistics_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	dateFrom, err := request.GetDateQueryParam(r, "date_from")
	if err != nil {
		rh.ErrorResponse(
			err,
			"invalid date_from",
		)
		return
	}

	dateTo, err := request.GetDateQueryParam(r, "date_to")
	if err != nil {
		rh.ErrorResponse(
			err,
			"invalid date_to",
		)
		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, userID, dateFrom, dateTo)
	if err != nil {
		rh.ErrorResponse(
			err,
			"failed to get statistics",
		)
		return
	}

	statisticsResponse := toResponse(statistics)

	rh.JSONResponse(
		http.StatusOK,
		statisticsResponse,
	)
}
