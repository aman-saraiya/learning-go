package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// anonymous function used to double the numbers
	transformed := transformNumbers(&numbers, func(number int) int { return number * 2 })

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// In programming in general or in Go, in cases where a function
// wants another function as a parameter value or where a function returns a function
// as a value, then some effort and time can be saved by using anonymous functions

// This feature allows us to define a function just in time when needed,
// instead of defining in advance.

// 	transformed := transformNumbers(&numbers, ???)

// consider this case where we might just want to use a particular tranform function
// only here and don't use it for any other purpose in the program
// In that case we can write an anonymous function in place of these ???
