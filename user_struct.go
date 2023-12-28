package main

import (
	"fmt"
	"time"
)

// defining a struct
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// NOTE here that the struct here is passed by value
func printUserDetails(appUser User) {
	fmt.Println("The first name of the user is: ", appUser.firstName)
	fmt.Println("The last name of the user is: ", appUser.lastName)
	fmt.Println("The birth date of the user is: ", appUser.birthDate)
	fmt.Println("The creation time of the user is: ", appUser.createdAt)
	fmt.Println()
}

func printUserDetailsUsingPointer(appUserPointer *User) {
	// ideally we need to dereference the pointers of other types before using
	// its value, but in case of structs we can directly use it as follows
	// to access the attributes and methods
	fmt.Println("The first name of the user is: ", appUserPointer.firstName)
	fmt.Println("The last name of the user is: ", appUserPointer.lastName)
	fmt.Println("The birth date of the user is: ", appUserPointer.birthDate)
	fmt.Println("The creation time of the user is: ", appUserPointer.createdAt)

	fmt.Println()

	// Alternative way, by dereferencing and using the pointer also works
	fmt.Println("The first name of the user is: ", (*appUserPointer).firstName)
	fmt.Println("The last name of the user is: ", (*appUserPointer).lastName)
	fmt.Println("The birth date of the user is: ", (*appUserPointer).birthDate)
	fmt.Println("The creation time of the user is: ", (*appUserPointer).createdAt)
}

// The above functions are standalone functions are not methods linked
// to the User struct

// methods are linked to struct by specifying the receiver parameter
// receiver parameter is a special kind of parameter that is not
// specfied in the parameter list but is specified in front of the
// function name to turn it into a method of that received type

func (u User) displayUserMethod() {
	fmt.Println("User's first name: ", u.firstName)
	fmt.Println("User's last name: ", u.lastName)
	fmt.Println("User's date of birth: ", u.birthDate)
	fmt.Println("User creation time: ", u.createdAt)
	fmt.Println()
}

// NOTE: The receiver parameter also works just like a regular paramter
// and it is passed by value, so to modify any field value we need to
// use pointer

func (u *User) modifyMe() {
	u.firstName = "Alex"
}

// defining constructor
func newUser(firstName, lastName, birthDate string) User {
	return User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}
}

func newUserPointer(firstName, lastName, birthDate string) *User {
	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}
}
