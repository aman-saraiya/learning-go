package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64

	revenue, err := getUserInput("Enter your revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err = getUserInput("Enter your tax rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err = getUserInput("Enter your expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	calculateFinancials(revenue, expenses, taxRate)
}

func calculateFinancials(revenue, expenses, taxRate float64) {
	earningsBeforeTax := revenue - expenses
	fmt.Println("Your Earnings Before Tax (EBT): ", earningsBeforeTax)
	earningsAfterTax := revenue*(1-taxRate/100) - expenses
	fmt.Print("Your Earnings After Tax (Profit): ", earningsAfterTax)
	result := fmt.Sprintf("EBT %.2f, Profit: %.2f", earningsBeforeTax, earningsAfterTax)
	os.WriteFile("financials.txt", []byte(result), 0644)
}

func getUserInput(prompt string) (float64, error) {
	fmt.Print(prompt)
	var input float64
	fmt.Scan(&input)
	if input < 0 {
		input = 0
		return input, errors.New("Invalid Input. Value must be greater than 0.")
	}
	return input, nil
}
