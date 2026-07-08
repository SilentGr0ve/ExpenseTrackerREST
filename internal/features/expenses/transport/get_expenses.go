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
