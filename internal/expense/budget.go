package expense

import "fmt"

type Budget struct {
	Month  int
	Amount float64
}

func NewBudget(month int, amount float64) *Budget {
	return &Budget{
		Month:  month,
		Amount: amount,
	}
}

func BudgetMonth(month int, amount float64) error {
	//read all budgets
	budgets, err := ReadBudgetData()

	if err != nil {
		return err
	}

	for i, budget := range budgets {
		if budget.Month == month {
			budgets[i].Amount = amount
			fmt.Printf("Budget for month %d updated to %.2f\n", month, amount)
			return WriteBudgetData(budgets)
		}
	}

	newBudget := NewBudget(month, amount)
	budgets = append(budgets, *newBudget)
	fmt.Printf("Budget for month %d set to %.2f\n", month, amount)
	return WriteBudgetData(budgets)
}

func GetMonthlyBudget(month int) (float64, error) {
	budgets, err := ReadBudgetData()

	if err != nil {
		return 0, err
	}

	for _, budget := range budgets {
		if budget.Month == month {
			return budget.Amount, nil
		}
	}

	return 0, nil
}
