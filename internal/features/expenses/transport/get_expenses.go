package expenses_transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
)

// GetExpenses godoc
// @Summary Get user expenses
// @Description Return a list of expenses for the authorized user with optional filters
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param category_id query string false "Filter by category UUID"
// @Param date_from query string false "Filter by start date (YYYY-MM-DD)"
// @Param date_to query string false "Filter by end date (YYYY-MM-DD)"
// @Param limit query int false "Limit number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} ExpenseResponse "Expense list returned"
// @Failure 400 {object} response.ErrorResponse "Invalid query parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /expenses [get]
func (h *ExpensesHTTPHandler) GetExpenses(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(
			core_errors.ErrUnauthorized,
			"unauthorized",
		)
		return
	}

	categoryID, dateFrom, dateTo, limit, offset, err := getParams(r)
	if err != nil {
		rh.ErrorResponse(
			err,
			"invalid query parameter",
		)
		return
	}

	expensesQuery := domain.ExpensesQuery{
		CategoryID: categoryID,
		DateFrom:   dateFrom,
		DateTo:     dateTo,
		Limit:      limit,
		Offset:     offset,
	}

	expenses, err := h.expensesService.GetExpenses(ctx, userID, expensesQuery)
	if err != nil {
		rh.ErrorResponse(
			err,
			"failed to get categories",
		)
		return
	}

	expensesResponse := ToDTOFromDomain(expenses)

	rh.JSONResponse(
		http.StatusOK,
		expensesResponse,
	)

}

func getParams(r *http.Request) (*uuid.UUID, *time.Time, *time.Time, *int, *int, error) {
	categoryID, err := request.GetUUIDQueryParam(r, "category_id")
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("invalid category_id")
	}

	dateFrom, err := request.GetDateQueryParam(r, "date_from")
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("invalid date_from")
	}

	dateTo, err := request.GetDateQueryParam(r, "date_to")
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("invalid date_to")
	}

	limit, err := request.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("invalid limit")
	}

	offset, err := request.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("invalid offset")
	}

	return categoryID, dateFrom, dateTo, limit, offset, nil
}
