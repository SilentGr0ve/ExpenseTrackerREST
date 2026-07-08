package statistics_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *StatisticsRepository) GetStatistics(ctx context.Context, userID uuid.UUID, dateFrom, dateTo *time.Time) (domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	period := domain.Period{
		From: dateFrom,
		To:   dateTo,
	}

	summaryQuery := `
	SELECT 
		COUNT(*) as total_count,
		COALESCE(SUM(amount), 0) as total_amount,
		COALESCE(AVG(amount), 0) as average_amount
	FROM tracker.expenses
	WHERE user_id=$1
	`

	argsSummary := []any{userID}
	argsSummaryIndex := 2

	if dateFrom != nil {
		summaryQuery += fmt.Sprintf(" AND expense_date >= $%d", argsSummaryIndex)
		argsSummary = append(argsSummary, dateFrom)
		argsSummaryIndex++
	}

	if dateTo != nil {
		summaryQuery += fmt.Sprintf(" AND expense_date <= $%d", argsSummaryIndex)
		argsSummary = append(argsSummary, dateTo)
		argsSummaryIndex++
	}

	row := r.pool.QueryRow(
		ctx,
		summaryQuery,
		argsSummary...,
	)

	summaryQuery += fmt.Sprintf(";")

	var summary domain.Summary

	if err := row.Scan(
		&summary.TotalExpenses,
		&summary.TotalAmount,
		&summary.AverageAmount,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Statistics{}, fmt.Errorf("user with id='%s' do not has expenses: %w", userID, core_errors.ErrNotFound)
		}
		return domain.Statistics{}, fmt.Errorf("scan error: %w", err)
	}

	byCategoryQuery := `
	SELECT 
	    c.id,
	    c.name, 
	    SUM(e.amount) as total_amount,
	    COUNT(*) as count
	FROM tracker.expenses e 
	JOIN tracker.categories c ON e.category_id = c.id
	WHERE e.user_id=$1
	`

	argsByCategory := []any{userID}
	argsByCategoryIndex := 2

	if dateFrom != nil {
		byCategoryQuery += fmt.Sprintf(" AND expense_date >= $%d", argsByCategoryIndex)
		argsByCategory = append(argsByCategory, dateFrom)
		argsByCategoryIndex++
	}

	if dateTo != nil {
		byCategoryQuery += fmt.Sprintf(" AND expense_date <= $%d", argsByCategoryIndex)
		argsByCategory = append(argsByCategory, dateTo)
		argsByCategoryIndex++
	}

	byCategoryQuery += fmt.Sprintf(" GROUP BY c.id, c.name ORDER BY total_amount DESC;")

	rows, err := r.pool.Query(
		ctx,
		byCategoryQuery,
		argsByCategory...,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get by category summary: %w", err)
	}
	defer rows.Close()

	var byCategorySummary []domain.CategorySummary

	for rows.Next() {
		var categorySummary domain.CategorySummary
		if err := rows.Scan(
			&categorySummary.CategoryID,
			&categorySummary.CategoryName,
			&categorySummary.TotalAmount,
			&categorySummary.ExpensesCount,
		); err != nil {
			return domain.Statistics{}, fmt.Errorf("scan category summary: %w", err)
		}
		byCategorySummary = append(byCategorySummary, categorySummary)
	}

	if err := rows.Err(); err != nil {
		return domain.Statistics{}, fmt.Errorf("rows error: %w", err)
	}

	statistics := domain.Statistics{
		Period:     period,
		Summary:    summary,
		ByCategory: byCategorySummary,
	}

	return statistics, nil
}
