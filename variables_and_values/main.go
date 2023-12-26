package main

import (
	"fmt"
	"math"
)

func main() {
	// declaring a variable; type inferred automatically
	var rateOfReturn = 4.5

	// declaring a variable; type specified explicitly
	var investmentAmount float64 = 1000

	// shorthand syntax for declaration and initialization; type inferred automatically
	yearsOfInvestment := 10

	// we cannot specify the type while using the short hand syntax
	// This will give error
	/*
		yearsOfInvestment float64 := 10
	*/

	// https://pkg.go.dev/math#Pow
	// math.Pow takes float64 as input
	// type conversion used for yearsOfInvestment var
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	var finalValue = investmentAmount * math.Pow(1+rateOfReturn/100, float64(yearsOfInvestment))
	fmt.Print("Final Value: ", finalValue, "\n")

	// Declaring multiple variables in a single statement
	var variableX, variableY = 100, 10 // both types inferred automatically
	fmt.Print(variableX, "\n")         // int
	fmt.Print(variableY, "\n")         // int

	var variableA, variableB = 100, 2.5 // both types inferred automatically
	fmt.Print(variableA, "\n")          // int
	fmt.Print(variableB, "\n")          // float64

	variableP, variableQ := "This is a string", 125.53 //shorthand syntax
	fmt.Print(variableP, "\n")                         // string
	fmt.Print(variableQ, "\n")                         // float64

	var variableM, variableN float64 = 100, 20.30 // common type explictly assigned for both variables
	fmt.Print(variableM, "\n")                    // float64
	fmt.Print(variableN, "\n")                    // float64

	/*
		var variableE float64, variableF string = 10.45, "This is a string"
	*/
	//	This will give error, we could explicitly specify only a single type with multiple declarations
	//	All the variables declared will be of type mentioned at the end of the statement
}
