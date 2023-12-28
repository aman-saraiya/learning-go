package main

import "fmt"

func workingWithLoops() {
	// We only have for loops in go
	// no while and do while loops like other languages
	// but for loop in go is pretty flexible

	for i := 0; i < 3; i++ { // initialization; check; increment
		workingWithConditionals()
	}
}

// if we skip all the initialization; check and increment we get an infinite loop
func workingWithInfiniteLoops() {
	for {
		workingWithConditionals()
	}
}

func loopWithBreakAndContinue() {
	for i := 1; ; i++ {
		if i < 10 {
			fmt.Println(i)
			continue // skips further code execution and starts the next loop iteration
		}
		fmt.Println("Breaking out of loop")
		break // breaks out of the loop
	}
}

/*
conditional for loops

for someCondition{
	...
}
Till the time someCondition yields true the loop continues
*/

func conditionalForLoop() {
	userInput := 10
	for userInput != 0 {
		fmt.Print("Enter new input: ")
		fmt.Scan(&userInput)
	}
}
