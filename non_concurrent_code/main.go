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
	greet("Nice to meet you")
	greet("How are you?")
	slowGreet("How ... are ... you ...?")
	greet("I hope you're having a great day")
}

// here we have the last greet function that doesn't depend on the slowGreet but its
// execution is blocked due to the slowGreet function, so here we could use the
// Goroutines to speed up the execution
// And therefore we can run them in parallel simply by adding a 'go' keyword in front of them

// Next: basic_concurrent_code/main.go
