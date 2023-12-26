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

	var finalValue = investmentAmount * math.Pow(1+rateOfReturn/100, float64(yearsOfInvestment))
	fmt.Print("Final Value: ", finalValue)
}
