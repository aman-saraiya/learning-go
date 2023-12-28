package main

import (
	"fmt"
	"time"
)

func greet(greeting string) {
	fmt.Println(greeting)
}

func slowGreet(greeting string) {
	time.Sleep(3 * time.Second) // simulating a slow time taking task
	fmt.Println(greeting)
}

func main() {
	go greet("Nice to meet you")
	go greet("How are you?")
	go slowGreet("How ... are ... you ...?")
	go greet("I hope you're having a great day")
}

/*
	By adding this 'go' keyword we tell go that we want to run this functions as go
	routines which in the end simply means they still will be executed but they will run
	in parallel.

	If we run this now, we see that it finished instantly but we don't really see any output!?
	We don't see anything on the console.

	This is because the go keyword just dispatches these functions to be executed in parallel
	and the dispatching is very quick so before any of the functions execute the main function
	reaches it's end and is done. And when the main function is done the go program is done.
	So, we don't see any output on the console.
*/

// To get the output we need to use channels
// Go to: concurrent_code_with_channels
