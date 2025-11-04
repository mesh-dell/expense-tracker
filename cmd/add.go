package cmd

import "github.com/mesh-dell/expense-tracker/internal/expense"

func AddCommand(description, category string, amount float64) error {
	return expense.AddExpense(amount, description, category)
}
