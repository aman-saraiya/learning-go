package main

import "fmt"

// Taking pointer as an parameter
func is_adult(age *int) bool {
	if *age >= 18 {
		return true
	} else {
		return false
	}
}

func main() {
	// declaring and initializing pointers
	age := 32
	agePointer := &age // type of agePointer is '*int'
	fmt.Println("Age is", age)
	fmt.Println("Address of Age variable is", &age)
	fmt.Println("AgePointer is", agePointer)

	fmt.Println("Dereferencing pointer to get values", *agePointer)

	// passing pointer as an argument
	fmt.Println("Is adult:", is_adult(agePointer))

	// passing address as an argument
	fmt.Println("Is adult:", is_adult(&age))
}
