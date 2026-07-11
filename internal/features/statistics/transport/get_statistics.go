package statistics_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

// GetStatistics godoc
// @Summary Get expense statistics
// @Description Returns aggregated expense statistics for the authorized user, optionally filtered by date range
// @Tags statistics
// @Produce json
// @Security BearerAuth
// @Param date_from query string false "Start date (YYYY-MM-DD)"
// @Param date_to query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} StatisticsResponse "Statistics returned"
// @Failure 400 {object} response.ErrorResponse "Invalid query parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /statistics [get]
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
