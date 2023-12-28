package main

import (
	"fmt"
)

func workingWithSwitch() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int = -1
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	fmt.Println("Selected:", choice)

	// if we have different code segment to be executed for different
	// values we could use switch statement instead of many if, else if blocks
	switch choice {
	case 1:
		fmt.Println("Your Balance:", accountBalance)
	case 2:
		var depositAmount float64
		fmt.Print("Deposit Amount: ")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid Deposit Amount")
			return
		}

		// equivalent to accountBalance = accountBalance + depositAmount
		accountBalance += depositAmount
		fmt.Println("New Balance:", accountBalance)
	case 3:
		var withdrawAmount float64
		fmt.Print("Deposit Amount: ")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 || withdrawAmount > accountBalance {
			fmt.Println("Withdraw amount invalid")
			return
			// main function doesn't return anything
			// but we can use the return statement anywhere in the function to
			// stop the function execution at that point
		}

		accountBalance -= withdrawAmount
		fmt.Println("New Balance:", withdrawAmount)
	default:
		fmt.Println("GoodBye!")
		break
	}
	// Note here we don't need to add break statements in each case, like we need in C/C++
}

// break keyword in switch statement has a special meaning
// here it no longer breaks out of the closest sorrounding loop but
// breaks out of the switch statement. But we don't generally need this feature.
