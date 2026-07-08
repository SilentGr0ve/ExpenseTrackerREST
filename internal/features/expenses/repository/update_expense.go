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

func (r *ExpensesRepository) UpdateExpense(ctx context.Context, userID uuid.UUID, expense domain.Expense) (domain.Expense, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	UPDATE tracker.expenses
	SET category_id=$1, amount=$2, description=$3,expense_date=$4, updated_at=NOW(),version=version+1
	WHERE id=$5 AND user_id=$6 AND version=$7
	RETURNING id, user_id, category_id,version,amount,description,expense_date,created_at,updated_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		expense.CategoryID,
		expense.Amount,
		expense.Description,
		expense.ExpenseDate,
		expense.ID,
		userID,
		expense.Version,
	)

	var patchedExpense domain.Expense

	if err := row.Scan(
		&patchedExpense.ID,
		&patchedExpense.UserID,
		&patchedExpense.CategoryID,
		&patchedExpense.Version,
		&patchedExpense.Amount,
		&patchedExpense.Description,
		&patchedExpense.ExpenseDate,
		&patchedExpense.CreatedAt,
		&patchedExpense.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Expense{}, core_errors.ErrConflict
		}
		return domain.Expense{}, fmt.Errorf("scan updated expense: %w", err)
	}

	return patchedExpense, nil

}
