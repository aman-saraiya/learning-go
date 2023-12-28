package main

import "fmt"

type CustomString string

// Could be thought of as go under the hood creates a copy of string type with
// name CustomString and we can then extend it
func (CustomString) log() {
	fmt.Println("This is a method added to the CustomString type.")
}

// Cannot do this
/*
func (string) log() {

}
*/
