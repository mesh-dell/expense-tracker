package expense

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func ReadExpensesData() ([]Expense, error) {
	filePath := GetExpensesFilePath()

	if !FileExists(filePath) {
		fmt.Println("Creating new expenses file.....")
		_, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("error creating file %s", err)
		}
		data := []byte("[]")
		os.WriteFile(filePath, data, os.ModeAppend.Perm())
		return []Expense{}, nil
	}

	//read file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file")
	}

	defer file.Close()
	var expenses []Expense
	err = json.NewDecoder(file).Decode(&expenses)

	if err != nil {
		fmt.Println("error decoding file")
		return nil, err
	}

	return expenses, nil
}

func WriteExpensesData(expenses []Expense) error {
	filePath := GetExpensesFilePath()
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	err = json.NewEncoder(file).Encode(expenses)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}
	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func GetExpensesFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	filePath := path.Join(cwd, "expenses.json")
	return filePath
}

func GetBudgetPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	filePath := path.Join(cwd, "budget.json")
	return filePath
}

func ReadBudgetData() ([]Budget, error) {
	filePath := GetBudgetPath()

	if !FileExists(filePath) {
		fmt.Println("Creating new Budget file......")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
		if err != nil {
			fmt.Println("Error creating file")
			return []Budget{}, err
		}
		defer file.Close()
		return nil, nil
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error creating file")
		return nil, err
	}

	defer file.Close()

	var budget []Budget
	err = json.NewDecoder(file).Decode(&budget)

	if err != nil {
		fmt.Println("Error decoding json")
		return nil, err
	}

	return budget, nil
}

func WriteBudgetData(budget []Budget) error {
	filepath := GetBudgetPath()
	file, err := os.Create(filepath)

	if err != nil {
		fmt.Println("Error creating file")
		return err
	}

	err = json.NewEncoder(file).Encode(budget)

	if err != nil {
		fmt.Println("Error encoding json")
		return err
	}

	return nil
}
