package expense

import (
	"fmt"
	"strings"
	"time"
)

type Expense struct {
	Id          int
	Description string
	Amount      float64
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewExpense(id int, description string, amount float64, category string) *Expense {
	expense := Expense{
		Id:          id,
		Description: description,
		Amount:      amount,
		Category:    category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &expense
}

func AddExpense(amount float64, description string, category string) error {
	expenses, err := ReadExpensesData()

	if err != nil {
		return err
	}

	var newExpenseId int

	if len(expenses) == 0 {
		newExpenseId = 1
	} else {
		var lastExpense = expenses[len(expenses)-1]
		newExpenseId = lastExpense.Id + 1
	}

	//check if budget exceeded
	thisMonth := time.Now().Month()
	thisMonthBudget, err := GetMonthlyBudget(int(thisMonth))

	if err != nil {
		return err
	}

	newExpense := NewExpense(newExpenseId, description, amount, category)
	expenses = append(expenses, *newExpense)

	thisMonthExpenses := 0.0

	for _, expense := range expenses {
		if expense.CreatedAt.Month() == thisMonth {
			thisMonthExpenses += expense.Amount
		}
	}

	if thisMonthBudget != 0 && thisMonthExpenses > thisMonthBudget {
		fmt.Printf("Warning: You have exceeded your budget for this month. Budget: %.2f, Expenses: %.2f\n", thisMonthBudget, thisMonthExpenses)
	}

	fmt.Printf("Expense Added: %s, Amount: %.2f, Category: %s\n", description, amount, category)
	return WriteExpensesData(expenses)
}

func ListExpenses(category string) error {
	expenses, err := ReadExpensesData()

	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("You have no expenses")
	}

	var filteredExpenses []Expense

	for _, expense := range expenses {
		if category == "all" || strings.EqualFold(category, expense.Category) {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}

	if len(filteredExpenses) == 0 {
		fmt.Println("No expenses found for this category")
		return nil
	}

	for _, expense := range filteredExpenses {
		fmt.Printf("ID: %d %s %.2f %s\n", expense.Id, expense.Description, expense.Amount, expense.Category)
	}
	return nil
}

func DeleteExpense(id int) error {
	expenses, err := ReadExpensesData()

	if err != nil {
		return err
	}

	var updatedExpenses []Expense
	for _, expense := range expenses {
		if expense.Id != id {
			updatedExpenses = append(updatedExpenses, expense)
		}
	}

	if len(updatedExpenses) == len(expenses) {
		return fmt.Errorf("expense id %d not found", id)
	}

	fmt.Println("Expense ID:", id, "deleted successfully")
	return WriteExpensesData(updatedExpenses)
}

func SummaryExpenses(month int) error {
	expenses, err := ReadExpensesData()

	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}

	var total float64
	if month == 0 {
		for _, expense := range expenses {
			total += expense.Amount
		}
	} else {
		for _, expense := range expenses {
			if expense.CreatedAt.Month() == time.Month(month) {
				total += expense.Amount
			}
		}
	}

	fmt.Printf("Total expenses: %.2f\n", total)
	return nil
}

func UpdateExpense(id int, amount float64, description, category string) error {
	expenses, err := ReadExpensesData()

	if err != nil {
		return err
	}

	var found bool
	var updatedExpenses []Expense

	for _, expense := range expenses {
		if expense.Id == id {
			// found
			found = true
			expense.Amount = amount
			expense.Description = description
			expense.Category = category
			expense.UpdatedAt = time.Now()
		}
		updatedExpenses = append(updatedExpenses, expense)
	}

	if !found {
		return fmt.Errorf("Expense of id:%d not found", id)
	}

	return WriteExpensesData(updatedExpenses)
}
