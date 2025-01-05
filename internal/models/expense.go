package models

import (
	"time"
)

type Expense struct {
	ID           int64      `json:"id"`
	UserID       int64      `json:"user_id"`
	Amount       int        `json:"amount" validate:"required,gt=0"`
	Date_expense string     `json:"date_expense" validate:"required,datetime=2006-01-02"`
	Category     Category   `json:"category" validate:"required"`
	Description  string     `json:"description" validate:"required,min=3,max=500"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type ExpenseRequest struct {
	Data *Expense `json:"data"`
}

type ExpenseResponse struct {
	Data *Expense `json:"data"`
}

type AllExpenses struct {
	Expenses []Expense `json:"expenses"`
}

type ExpensesResponse struct {
	Data *AllExpenses `json:"data"`
}
