package main

import "fmt"

func main() {
	//working with type switches
	workingWithTypeSwitches(12)
	workingWithTypeSwitches(9.55)
	workingWithTypeSwitches(myNote)
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
