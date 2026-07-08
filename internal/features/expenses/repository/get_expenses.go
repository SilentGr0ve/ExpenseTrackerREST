package expenses_repository

import (
	"context"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

func (r *ExpensesRepository) GetExpenses(ctx context.Context, userID uuid.UUID, expenseQuery domain.ExpensesQuery) ([]domain.Expense, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
    SELECT id, user_id, category_id, version, amount, description, expense_date, created_at, updated_at
    FROM tracker.expenses
    WHERE user_id = $1
    `

	args := []any{userID}
	argsIndex := 2

	if expenseQuery.CategoryID != nil {
		query += fmt.Sprintf(" AND category_id=$%d", argsIndex)
		args = append(args, *expenseQuery.CategoryID)
		argsIndex++
	}

	if expenseQuery.DateFrom != nil {
		query += fmt.Sprintf(" AND expense_date >= $%d", argsIndex)
		args = append(args, *expenseQuery.DateFrom)
		argsIndex++
	}

	if expenseQuery.DateTo != nil {
		query += fmt.Sprintf(" AND expense_date <= $%d", argsIndex)
		args = append(args, *expenseQuery.DateTo)
		argsIndex++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d;", argsIndex, argsIndex+1)
	args = append(args, *expenseQuery.Limit, *expenseQuery.Offset)

	rows, err := r.pool.Query(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, fmt.Errorf("get expenses: %w", err)
	}
	defer rows.Close()

	var expenses []domain.Expense

	for rows.Next() {
		var expense domain.Expense
		if err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.Version,
			&expense.Amount,
			&expense.Description,
			&expense.ExpenseDate,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan expense: %w", err)
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return expenses, nil
}
