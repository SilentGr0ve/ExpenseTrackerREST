package expenses_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *ExpensesService) DeleteExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) error {
	err := s.expensesRepository.DeleteExpense(ctx, userID, expenseID)
	if err != nil {
		fmt.Errorf("delete expense from repository: %w", err)
	}

	return nil
}
