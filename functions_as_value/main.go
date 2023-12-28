package main

import "fmt"

/*
Functions, are first class values in go.
We can use functions themselves as parameter values for other functions.
*/
func main() {
	numbers := []int{1, 2, 3, 6}
	doubledNumbers := doubleNumbers(&numbers)
	fmt.Println(doubledNumbers)

	// passing functions as arguments
	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))

	// functions as return value
	fmt.Println(transformNumbers(&numbers, getTransformerFunction("double")))
	fmt.Println(transformNumbers(&numbers, getTransformerFunction("triple")))
}

// func doubleNumbers(numbers *[]int) []int {
// 	doubledNumbers := []int{}
// 	for _, value := range *numbers {
// 		doubledNumbers = append(doubledNumbers, value*2)
// 	}
// 	return doubledNumbers
// }

// now let's say we want a double numbers function that doesn't work only on
// slices but instead we want a general function that doubles the integer itself

func double(number int) int {
	return number * 2
}

// Now we could use this double function inside the doubleNumber function as
func doubleNumbers(numbers *[]int) []int {
	doubledNumbers := []int{}
	for _, value := range *numbers {
		doubledNumbers = append(doubledNumbers, double(value))
	}
	return doubledNumbers
}

// but here we're passing the result of calling that function to append
// still we're not using function as parameter anywhere

// Now let's say we have a triple function as well
func triple(number int) int {
	return number * 3
}

//now let's say we want a generic transform function that processes the input based on the
// function we pass as a parameter

// understanding the syntax here:
// transform func()
// transform is the name of the parameter, we could use anything
// func is the type
// in paraentheses we don't define the parameters but just the types
// of the parameters the function should have and the last part to define
// the return type of the function
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}
	for _, val := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(val))
	}
	return transformedNumbers
}

// function (int) int types like this could get rather long depending on the
// input params and return values. So we could use alias (custom types) for this like

type transformFunctionType func(int) int

// getting function as a return value

func getTransformerFunction(multiple string) func(int) int {
	if multiple == "double" {
		return double
	} else if multiple == "triple" {
		return triple
	} else {
		return nil
	}
}
