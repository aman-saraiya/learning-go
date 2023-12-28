package main

import "fmt"

func main() {
	doubleTransformer := createTransformers(2)
	fmt.Println(doubleTransformer(15))

	tripleTransformer := createTransformers(3)
	fmt.Println(tripleTransformer(15))
}

// Note here that factor is not a parameter of the anonymous function but of createTransformer
// But because of scoping and because of the fact that we can use variables and parameters act
// like variables and parameters act like variables, in lower level nested scopes means that we
// can use factor which belongs to the scope of create transformer inside the anonymous function
// inside the createTransformer
// Now we can use createTransformer as a so called factory function because createTransformer now
// is a function that produces other function with different configurations, and this means we can now
// easily create different number transformers.
func createTransformers(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Here the anonymous function in createTransformers is a closure.
// If we use variable from the scope in which the function is created
// then the value of that outer scope variable is locked in the function created inside.

// so here doubleTransformer := createTransformers(2) and 	tripleTransformer := createTransformers(3)
// we are calling createTransfomer two times first to create doubleTransformer and second to create
// tripleTransformer but creating of second 'tripleTransformer' doesn't change the first one.
// The value 2 is locked in the doubleTransformer and similarly for value 3 in tripleTransformer.
