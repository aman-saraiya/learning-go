package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	fmt.Println("Profit Calculator")
	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	earningsBeforeTax := revenue - expenses
	fmt.Println("Your Earnings Before Tax (EBT): ", earningsBeforeTax)
	earningsAfterTax := revenue*(1-taxRate/100) - expenses
	fmt.Print("Your Earnings After Tax (Profit): ", earningsAfterTax)
}
