package expenses_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
)

func (s *ExpensesService) PatchExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID, patch domain.ExpensePatch) (domain.Expense, error) {
	currentExpense, err := s.expensesRepository.GetExpense(ctx, userID, expenseID)
	if err != nil {
		return domain.Expense{}, fmt.Errorf("get expense from repository: %w", err)
	}

	if patch.CategoryID != nil {
		currentExpense.CategoryID = *patch.CategoryID
	}

	if patch.Amount != nil {
		currentExpense.Amount = *patch.Amount
	}

	if patch.Description != nil {
		currentExpense.Description = patch.Description
	}

	if patch.ExpenseDate != nil {
		currentExpense.ExpenseDate = *patch.ExpenseDate
	}

	expense, err := s.expensesRepository.UpdateExpense(ctx, userID, currentExpense)
	if err != nil {
		if errors.Is(err, core_errors.ErrConflict) {
			return domain.Expense{}, fmt.Errorf("expense modified earlier: %w", err)
		}
		return domain.Expense{}, fmt.Errorf("update expense from repository: %w", err)
	}

	return expense, nil
}
