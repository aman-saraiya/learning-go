package main

import (
	"fmt"
	"time"
)

func main() {
	var appUser User // declaring a struct

	//initializing struct
	appUser = User{
		firstName: "Foo",
		lastName:  "Bar",
		birthDate: "01/01/2001",
		createdAt: time.Now(),
	}
	fmt.Println(appUser.firstName) // accessing struct fields

	// if the values are assigned in the order they're declared we
	// can omit the keys as follows
	appUserOne := User{
		"Alpha",
		"Beta",
		"02/02/2002",
		time.Now(),
	}

	fmt.Println(appUserOne.firstName)

	// creating an empty struct
	appUserEmpty := User{}

	fmt.Println(appUserEmpty.firstName) // empty string

	// passing structs to functions
	/*
		printUserDetails(appUser)

		printUserDetailsUsingPointer(&appUserOne)
	*/

	// using struct methods
	appUser.displayUserMethod()

	appUser.modifyMe()

	appUser.displayUserMethod()

	appUserConstructed := newUser("Sine", "Tan", "03/03/2003")
	appUserConstructed.displayUserMethod()

	appUserConstructedPointer := newUserPointer("Cos", "Cot", "04/04/2004")
	appUserConstructedPointer.displayUserMethod()

	admin := NewAdmin("alpha@gmail.com", "alphapass")
	adminAnonUser := NewAdminAnon("beta@gmail.com", "betapass")

	// difference in accessing the fields depending on if the embedded
	// struct is named or anon
	fmt.Println(admin.user.firstName)
	// fmt.Println(admin.firstName) // ERROR
	fmt.Println(adminAnonUser.firstName)

	// declaring and using variable of the custom type
	var my_custom_string CustomString
	my_custom_string.log()
}
