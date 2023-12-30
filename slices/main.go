package main

import "fmt"

func main() {
	// Building dynamic lists with slices
	// prices := []float64{} // if we don't specify the number of items stored in the array,
	// Go will automatically create a slice for us and since the slice is always based on an
	// array it will also create an array for us behind the scenes but it will automatically
	// ditch that array and create a new array if our slice grows beyond the bounds of that
	// behind the scenes storage array.
	// we could also add some initial values
	prices := []float64{10.99, 11.10}

	// We can use this further similar to how we used slices earlier
	// We could build slices on top of this slice, we could select individual elements,
	fmt.Println(prices[0:1])

	// we could also assign new values to the slice elements
	prices[1] = 99.99
	fmt.Println(prices)

	// We cannot use indexes, that are not a part of the slice
	// prices[2] = 101.11 // ERROR
	// so we still can't access indexes that don't exist
	// but to access the dynamic capabilities of the slice, we need
	// to use the built in append function in go
	// This returns a brand new slice

	// It does not edit the original slice but instead what append does is,
	// we want to add the item to that slice and therefore also to the underlying array.
	// Now, ofcourse array has a fixed length. So what go will do in this case is it will
	// create a brand new array and add that element to the new array
	newPrices := append(prices, 85.38, 91.38)
	fmt.Println(newPrices)
	// NOTE here that the original slice prices is unaffected
	fmt.Println(prices)
	// Is there also a function to remove elements from a slice
	// kind off
	// but there's no function for doing that
	prices = prices[1:] // removing the first element

	// Unpacking list values
	// append can take multiple values and appends it to the slice.
	// sometimes we want to append a slice to another slice
	discountPrices := []float64{199.99, 20.91, 101.11}
	//merging discountPrices to prices
	prices = append(prices, discountPrices...)
	// This syntax is used to unpack the list values
	fmt.Println(prices)
}
