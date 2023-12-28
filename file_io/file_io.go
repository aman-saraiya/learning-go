package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeIntToFile(value int) {
	// creates the file if doesn't exist, else overwrites the existing file
	// the second param required by WriteFile is of type []byte, consider byte collection
	// we can convert a string to byte collection as []byte(<string>)
	balanceText := fmt.Sprint(value)                       // converting int to string
	os.WriteFile("balance.txt", []byte(balanceText), 0644) // converting string to []byte
	// last parameter is the permission that should be assigned to the file
	// if the file already exists the permissions are not changed

	// 0644 -> owner rw, group r, other r
}

func readIntFromFile() {
	bytes, _ := os.ReadFile("balance.txt") // returns []byte, error
	// _ is a special kind of variable name we can use in situations when we want
	// to tell go that we're aware of the value we're getting and we don't want to use it
	// if we use any other name we will get an error that variable in not used

	balanceText, _ := strconv.ParseFloat(string(bytes), 64)
	// we cannot directly convert string to float using float64(<string>).
	fmt.Println("Balance text read from file: ", balanceText)
}
