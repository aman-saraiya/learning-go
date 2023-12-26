package main

import "fmt"

func main() {
	fmt.Print("This is a simple function to print on console", "\n")

	// we can pass any number of arguments to fmt.Print
	// They are printed without any space between them
	fmt.Print("This is First Argument", "This is second Argument", "\n")

	// To add a new line at the end
	fmt.Println("This prints a new line at the end")

	// we can pass any number of arguments to fmt.Print
	// They are printed without any space between them
	fmt.Println("This is first Argument.", "This is second Argument.")

	var name string
	// if we don't initialize a variable we cannot use short hand syntax
	// we need to use var and explicilty mention the type of the var as above

	fmt.Print("Enter your name: ")

	// Reading input from user
	fmt.Scan(&name)
	fmt.Println("Welcome to the Go community", name)
	// NOTE: We cannot easily fetch multi word input with fmt.Scan

}
