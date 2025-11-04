package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
}

func loadExpense() []Expense {
	file, err := os.ReadFile("expenses.json")

	if err != nil {
		return []Expense{}
	}

	var expenses []Expense
	json.Unmarshal(file, &expenses)

	return expenses
}

func saveExpense() {
	data, err := json.MarshalIndent(expenses, "", "  ")

	if err != nil {
		fmt.Println("Error while encoding:", err)
		return
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		fmt.Println("Error while writing:", err)
		return
	}
}

var expenses []Expense

func addExpense(rest []string) {
	var lasti int = len(expenses) - 1
	var idof int

	if lasti < 0 {
		idof = 0
	} else {
		idof = expenses[lasti].ID
	}
	idof++

	now := time.Now()
	date := now.Format("2006-01-02")

	var descriptionM string
	var amountM string

	for j, c := range rest {
		if len(c) > 3 && c[2:] == "description" {
			descriptionM = rest[j+1]
		} else if len(c) > 3 && c[2:] == "amount" {
			num, err := strconv.ParseFloat(rest[j+1], 64)

			if err != nil {
				fmt.Print("Not valid amount \n")
				return
			}

			if j+1 < len(rest) && num >= 0 {
				amountM = rest[j+1]
			}

		}
	}

	expense := Expense{
		ID:          idof,
		Date:        date,
		Description: descriptionM,
		Amount:      amountM,
	}

	expenses = append(expenses, expense)
	fmt.Print("Added succesfully \n")
	saveExpense()
}

func listExpense() {
	fmt.Printf("%-5s %-12s %-25s %-10s\n", "ID", "Date", "Description", "Amount")
	fmt.Println("---------------------------------------------------------------")
	for _, e := range expenses {
		fmt.Printf("%-5d %-12s %-25s %-10s\n", e.ID, e.Date, e.Description, e.Amount)
	}
	fmt.Println("---------------------------------------------------------------")
}

func deleteExpense(rest []string) {
	var index int = -1

	if len(rest) == 3 {
		id, err := strconv.Atoi(rest[2])
		if err != nil {
			fmt.Println("Conversion error:", err)
		}

		for i, v := range expenses {
			if v.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("Expense not found")
			return
		}
	} else {
		fmt.Print("Incorrect usage of delete")
	}

	expenses = append(expenses[:index], expenses[index+1:]...)
	fmt.Print("Expense deleted succesfully \n")
	saveExpense()
}

func summaryExpense(rest []string) {
	var total float64 = 0
	if len(rest) == 1 {
		for _, v := range expenses {
			num, err := strconv.ParseFloat(v.Amount, 64)
			if err != nil {
				fmt.Println("Conversion error:", err)
				return
			}
			total += num

		}
		fmt.Printf("Total expenses: $%.2f\n", total)
	} else if len(rest) == 3 {
		for _, v := range expenses {
			if len(v.Date) >= 7 && v.Date[5:7] == rest[2] {
				num, err := strconv.ParseFloat(v.Amount, 64)
				if err != nil {
					fmt.Println("Conversion error:", err)
					return
				}
				total += num
			}

		}
		fmt.Printf("Total expenses in the %s month: $%.2f\n", rest[2], total)
	}

}

func updateExpense(rest []string) {

	if len(expenses) <= 0 {
		print("No expense")
		return
	}
	id, err := strconv.Atoi(rest[2])
	if err != nil {
		fmt.Println("Conversion error:", err)
	}
	var index int = -1

	for i, v := range expenses {
		if v.ID == id {
			index = i
			break
		}
	}

	for j, c := range rest {
		if len(c) > 3 && c[2:] == "description" {
			expenses[index].Description = rest[j+1]
		} else if len(c) > 3 && c[2:] == "amount" {
			if j+1 < len(rest) {
				expenses[index].Amount = rest[j+1]
			}

		}
	}
	fmt.Printf("updated succesfully \n")
	saveExpense()
}

func main() {
	expenses = loadExpense()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		parts := strings.Fields(input)

		command := parts[0]
		rest := parts

		switch command {
		case "add":
			addExpense(rest)
		case "update":
			updateExpense(rest)
		case "delete":
			deleteExpense(rest)
		case "list":
			listExpense()
		case "summary":
			summaryExpense(rest)
		default:
			fmt.Println("Unknown command.")
		}

	}

}
