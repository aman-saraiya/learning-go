package main

import "fmt"

func main() {
	// here we're trying to read a file that doesn't exist.
	// This function still completes successfully and we get the balance value 0
	// So, in go when an error is encountered the app doesn't crash
	// os.ReadFile when it doesn't find a file, it generates an error but doesn't crash the app
	// In Go error handling works differently than other programming languages
	_, err := readIntFromFile()
	if err != nil {
		fmt.Println(err)
	}
}
