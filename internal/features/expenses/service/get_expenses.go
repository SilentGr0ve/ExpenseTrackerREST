package expenses_service

import (
	"context"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (s *ExpensesService) GetExpenses(ctx context.Context, userID uuid.UUID, query domain.ExpensesQuery) ([]domain.Expense, error) {

	limit := 20
	if query.Limit != nil {
		limit = *query.Limit
	}

	offset := 0
	if query.Offset != nil {
		offset = *query.Offset
	}

	query.Limit = &limit
	query.Offset = &offset

	if query.DateFrom != nil && query.DateTo != nil {
		if query.DateTo.Before(*query.DateFrom) {
			return nil, fmt.Errorf("date_to can't be before date_from: %w", core_errors.ErrInvalidArgument)
		}
	}

	if query.Limit != nil && *query.Limit < 0 || *query.Limit > 100 {
		return nil, fmt.Errorf("limit must be 1-100: %w", core_errors.ErrInvalidArgument)
	}

	if query.Offset != nil && *query.Offset < 0 {
		return nil, fmt.Errorf("offset must be >= 0: %w", core_errors.ErrInvalidArgument)
	}

	expenses, err := s.expensesRepository.GetExpenses(ctx, userID, query)
	if err != nil {
		return nil, fmt.Errorf("get expenses from repository: %w", err)
	}

	return expenses, nil

}
