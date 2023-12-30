package main

import "fmt"

func main() {
	sum := sumUp(10, 20, 30, 50)
	fmt.Println(sum)

	// passing slice to such variadic functions
	// splitting slice into parameter values
	nums := []int{1, 2, 4, 6}
	sum = sumUp(nums...)
	fmt.Println(sum)
}

// let's say we want a sumUp function that takes in variable number of
// parameters and returns the sum, although we could use slices to pass
// this numbers together to the function
// Sometimes though while writing code or when working on more complex programs
// we might not have a slice or an array of numbers available, we might want to
// call sumUp but instead of creating slice and adding an extra step we might like to
// pass in our list of numbers as simple parameter values

// we can do that using variadic functions

func sumUp(numbers ...int) int {
	sum := 0
	// Inside the function this numbers is still a slice
	// so the ... dots in the end collects that list of standalone
	// values and packages that all into a slice for us
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// func sumUpTemp(numbers ...int, temp int) int{} //NOTE THIS IS INVALID
// can only use ... with final parameter in list

// func sumUpTemp2(numbers ...int, temp string) // this also will not work

// func sumUpTemp3(temp int, numbers ...int) int {} // THIS IS VALID
