package main

import (
	"fmt" // use double quotes
)

var Temperature float64 = 98

// Simple function with two parameters
func outputText(name string, value float64) {
	fmt.Println(name, value)
}

// if the params are of same type we can define the type once at the end
func outputTextThreeParam(param1, param2 float64, text string) {
	fmt.Println(param1, param2, text)
}

// if the params are of same type we can define the type once at the end
func FunctionWithTwoReturnValues(param1, param2 float64) (float64, float64) {
	return param1, param2
}

// Alternative return value syntax for functions
func alternativeReturnValueSyntax(param1 string, param2 float64) (ret1 string, ret2 float64) { // these syntax declares ret1 and ret2 here
	ret1 = param1
	ret2 = param2 // we don't use := here as we don't want to declare + assign,
	// the variables are already declared in the function declaration syntax
	// we just directly use them in the function body
	return // this returns ret1, ret2
}

func main() {
	outputText("Value of gravitational acceleration (in m/s^2) is", 9.8)
	outputTextThreeParam(99.99, 9.9, "This are two random numbers.")
	fmt.Println(FunctionWithTwoReturnValues(10.10, 20.20))
	return_v1, return_v2 := FunctionWithTwoReturnValues(10.10, 20.20)
	fmt.Println("Return value 1:", return_v1)
	fmt.Println("Return value 2:", return_v2)

	fmt.Println(alternativeReturnValueSyntax("Age", 32))
}
