package main

import (
	"fmt"

	"github.com/aman-saraiya/learning-go/interfaces/note"
)

func main() {
	myNote, _ := note.New("Pay Electricity Bill", "Due Date: 30/01/2024")
	//working with type switches
	workingWithTypeSwitches(12)
	workingWithTypeSwitches(9.55)
	workingWithTypeSwitches(myNote)

	addApproachOne(2, 3)
	addApproachOne(6.5, 1.2)
	addApproachOne(myNote, myNote) // does nothing

	output1 := addUsingGenerics(2, 10)
	fmt.Println(output1)
	output2 := addUsingGenerics(2.4, 1.4)
	fmt.Println(output2)
}

// Function that can take any kind of value int, any struct, anything
func doSomething(input interface{}) {

}

// we can also use `any`, it is an alias for this interface{}
func doSomethingOne(input any) {

}

// Working with Type Switches
func workingWithTypeSwitches(input interface{}) {
	switch input.(type) { // special syntax that returns type
	case int:
		fmt.Println("Passed Input is int.")
	case float64:
		fmt.Println("Passed Input is float64.")
	case string:
		fmt.Println("Passed Input is string.")
	default:
		fmt.Println("Passed input is not a int, float or string.")
	}

	// note that input.(type) this syntax works only inside the
	// type switch statement.

	// but there's an alternative approach to extract type information
	// from values
	// we can pass a specific type name for which we're looking a value
	typedValue, ok := input.(int)
	if ok {
		fmt.Println("Entered input is an integer", typedValue)
	}
}

// Writing generic functions
// we want to write an add function for example with works with all int, float
// and string. If we use interface{} that will allow any types and not all types
// support '+' operation and hence we'll get an error

// approach 1
func addApproachOne(a, b interface{}) interface{} {
	aInt, aOk := a.(int)
	bInt, bOk := a.(int)
	if aOk && bOk {
		fmt.Println("Sum of a and b:", aInt+bInt)
		return aInt + bInt
	}

	aFloat, aOk := a.(float64)
	bFloat, bOk := b.(float64)
	if aOk && bOk {
		fmt.Println("Sum of a and b:", aFloat+bFloat)
		return aFloat + bFloat
	}

	return nil
}

// approach 2: Using generics
// T here is a type placeholder, we could use any name
// T <concrete types allowed in place of this placeholder>
func addUsingGenerics[T int | string | float64](a, b T) T {
	return a + b
}

// sample usage of multiple generic types
func usingMultipleGenericTyoes[T any, K any](x T, y K) T { return x }
