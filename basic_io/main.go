package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	// fmt.Scan doesn't stop reading until at least one non whitespace character is recorded
	// we cannot just hit enter to pass an empty value
	// This feature can be achieved using fmt.Scanln

	var optionalValue int
	fmt.Print("Enter value (optional): ")
	fmt.Scanln(&optionalValue)
	fmt.Println("Optional value:", optionalValue)
	// Here the fmt.Scanln isn't waiting for our input as it is
	// taking the new line left over from the previous input that is
	// not read by Scan

	// If we use Scanln at both places this works fine.
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Welcome to the Go community", name)
	fmt.Print("Enter value (optional): ")
	fmt.Scanln(&optionalValue)
	fmt.Println("Optional value:", optionalValue)

	// still both fmt.Scan and fmt.Scanln reads only single word/ no words
	// To read long form text we need to work through more complex steps
	text := getLongFormUserInput("Enter your address: ")
	fmt.Println(text)
}

func getLongFormUserInput(prompt string) string {
	fmt.Print(prompt)
	// We cannot use fmt.Scan or fmt.Scanln here as they don't really work with
	// long user input text, they work only with single value inputs
	/*
		fmt.Scan(&value)
	*/

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // \n is the byte where it should stop reading (delimiter)
	// Don't use double quotes here, we need to use single quotes because this is
	// actually a special value type in Go, called Rune, which is effectively a single
	// character

	if err != nil {
		return ""
	}

	// Note the read text still contains the delimiter. we need to process the string to remove it.
	// Can be done using strings package

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // this is required as in windows new line is \r\n
	return text
}
