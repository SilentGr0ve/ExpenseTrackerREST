package statistics_transport

import (
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StatisticsResponse struct {
	Period     PeriodResponse            `json:"period"`
	Summary    SummaryResponse           `json:"summary"`
	ByCategory []CategorySummaryResponse `json:"by_category"`
}

type PeriodResponse struct {
	From string `json:"date_from" example:"2026-01-01"`
	To   string `json:"date_to" example:"2026-12-31"`
}

type SummaryResponse struct {
	TotalAmount   decimal.Decimal `json:"total_amount" example:"9999.99"`
	TotalExpenses int             `json:"total_expenses" example:"10"`
	AverageAmount decimal.Decimal `json:"average_amount" example:"999.99"`
}

type CategorySummaryResponse struct {
	CategoryID    uuid.UUID       `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CategoryName  string          `json:"category_name" example:"Fastfood"`
	TotalAmount   decimal.Decimal `json:"total_amount" example:"999.99"`
	ExpensesCount int             `json:"expenses_count" example:"1"`
}

func toResponse(statistics domain.Statistics) StatisticsResponse {

	var fromStr, toStr string
	if statistics.Period.From != nil {
		fromStr = statistics.Period.From.Format("2006-01-02")
	}
	if statistics.Period.To != nil {
		toStr = statistics.Period.To.Format("2006-01-02")
	}
	periodResponse := PeriodResponse{
		From: fromStr,
		To:   toStr,
	}

	summaryResponse := SummaryResponse{
		TotalAmount:   statistics.Summary.TotalAmount,
		TotalExpenses: statistics.Summary.TotalExpenses,
		AverageAmount: statistics.Summary.AverageAmount,
	}

	var byCategorySummary []CategorySummaryResponse

	for _, category := range statistics.ByCategory {

		categorySummaryResponse := CategorySummaryResponse{
			CategoryID:    category.CategoryID,
			CategoryName:  category.CategoryName,
			TotalAmount:   category.TotalAmount,
			ExpensesCount: category.ExpensesCount,
		}

		byCategorySummary = append(byCategorySummary, categorySummaryResponse)
	}

	statisticsResponse := StatisticsResponse{
		Period:     periodResponse,
		Summary:    summaryResponse,
		ByCategory: byCategorySummary,
	}

	return statisticsResponse
}
