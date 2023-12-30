package main

import "fmt"

func main() {

	// Declaring and initializing array
	// we can create arrays of any types, including integers, strings, structs and arrays itself
	prices := [4]float64{10.99, 9.10, 15.12, 22.23}
	fmt.Println(prices)

	// Another way of creating arrays
	var productNames []string // DECLARATION
	// we could also add the length of array between []
	//like var productNames [4]string // Initialized to null values of the type
	productNames = []string{"CPU", "GPU", "Keyboard", "Screen"} // INITIALIZATION
	fmt.Println(productNames)

	// accessing individual values
	fmt.Println(productNames[2]) // 0 indexed
	// Error: index out of bounds
	/*
		fmt.Println(prices[5])
		fmt.Println(productNames[5])
	*/

	//assigning value
	productNames[1] = "Memory"
	fmt.Println(productNames)

	//Selecting parts of arrays with slices
	// [start, end)
	fmt.Println(productNames[0:3])
	slicedProductNames := productNames[0:3]
	fmt.Println(slicedProductNames)
	slicedProductNames[0] = "GPU"
	fmt.Println(slicedProductNames)
	fmt.Println(productNames)
	/*
	   If you simply assign an existing slice to a new variable,
	   the slice will not be duplicated. Both variables will refer to the exact same slice,
	   so any changes to the slice's value will be reflected in all its references.
	   This is also true when passing a slice to a function.
	*/

	// More ways of Selecting slices
	slice_one := productNames[:2]
	fmt.Println(slice_one)
	slice_two := productNames[2:]
	fmt.Println(slice_two)
	// Negative indexes works in some other programming languages but that doesn't work here
	// We also cannot pick a higher end index than the original array has

	// we can further slice a slice
	slice_a := productNames[0:3]
	slice_b := slice_a[0:1]
	fmt.Println(slice_a)
	fmt.Println(slice_b)

	/*
	   Slices are like a reference, like a window into an array, a bit like a pointer,
	   though it's a different concept
	   When we create an array like prices that array is stored in memory, when we create
	   a slice based on the array we get a window into the array, so to say. Therefore,
	   if we would modify an element in the slice, we would also modify the same element
	   in the original array.
	*/

	// Go also saves some metadata for our slices that could be useful to look into
	fmt.Println(slice_one)
	// len is a default go function, we don't need to import anything to use it
	// we can also call it on other value types like arrays, string etc
	fmt.Println(len(slice_one))
	// cap is also a default go function
	fmt.Println(cap(slice_one))
	fmt.Println(slice_two)
	fmt.Println(len(slice_two))
	fmt.Println(cap(slice_two))
	/*
	   Difference between len and cap?
	   len is the number of items in the slice or an array
	   cap is a bit more complex
	   It indicates how much a slice can extend towards the end
	   and this is important to understand about slices that you
	   can always select more towards the end of the array but not
	   towards the start
	   The capacity only counts towards the end of the array being sliced
	   but omits any elements that might have been filtered out before.
	*/

	fmt.Println("Understanding Capacity")
	featuredPrices := []int{10, 12, 15, 19}
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	// extending towards the right
	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices))
	fmt.Println(cap(highlightedPrices))
}
