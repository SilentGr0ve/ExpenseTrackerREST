package expenses_service

import (
	"context"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

func (s *ExpensesService) GetExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) (domain.Expense, error) {
	expense, err := s.expensesRepository.GetExpense(ctx, userID, expenseID)
	if err != nil {
		return domain.Expense{}, fmt.Errorf("get expense from repository: %w", err)
	}

	return expense, nil
}
