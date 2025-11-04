package cmd

import (
	"github.com/mesh-dell/expense-tracker/internal/expense"
)

func UpdateCommand(id int, amount float64, description, category string) error {
	return expense.UpdateExpense(id, amount, description, category)
}
