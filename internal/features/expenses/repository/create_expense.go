package expenses_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ExpensesRepository) CreateExpense(ctx context.Context, expense domain.Expense) (domain.Expense, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	query := `
	INSERT INTO tracker.expenses 
	    (id, user_id, category_id, version, amount, description, expense_date,created_at, updated_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id, user_id, category_id, version, amount, description, expense_date, created_at, updated_at;
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		expense.ID,
		expense.UserID,
		expense.CategoryID,
		expense.Version,
		expense.Amount,
		expense.Description,
		expense.ExpenseDate,
		expense.CreatedAt,
		expense.UpdatedAt,
	)

	var newExpense domain.Expense

	if err := row.Scan(
		&newExpense.ID,
		&newExpense.UserID,
		&newExpense.CategoryID,
		&newExpense.Version,
		&newExpense.Amount,
		&newExpense.Description,
		&newExpense.ExpenseDate,
		&newExpense.CreatedAt,
		&newExpense.UpdatedAt,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.Expense{}, core_errors.ErrConflict
		}
		return domain.Expense{}, fmt.Errorf("scan expense model: %w", err)
	}

	return newExpense, nil
}
