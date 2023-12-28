package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	revenue = getUserInput("Enter your revenue: ")
	taxRate = getUserInput("Enter your tax rate: ")
	expenses = getUserInput("Enter your expenses: ")

	calculateFinancials(revenue, expenses, taxRate)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64) {
	earningsBeforeTax := revenue - expenses
	fmt.Println("Your Earnings Before Tax (EBT): ", earningsBeforeTax)
	earningsAfterTax := revenue*(1-taxRate/100) - expenses
	fmt.Print("Your Earnings After Tax (Profit): ", earningsAfterTax)
	return earningsBeforeTax, earningsAfterTax
}

func getUserInput(prompt string) float64 {
	fmt.Print(prompt)
	var input float64
	fmt.Scan(&input)
	return input
}
