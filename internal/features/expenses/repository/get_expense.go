package expenses_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ExpensesRepository) GetExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) (domain.Expense, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	SELECT id, user_id, category_id, version, amount, description, expense_date, created_at, updated_at
	FROM tracker.expenses
	WHERE id=$1 AND user_id=$2;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		expenseID,
		userID,
	)

	var expense domain.Expense

	if err := row.Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Expense{}, fmt.Errorf("expense with id='%s': %w", expenseID, core_errors.ErrNotFound)
		}

		return domain.Expense{}, fmt.Errorf("scan error: %w", err)
	}

	return expense, nil

}
