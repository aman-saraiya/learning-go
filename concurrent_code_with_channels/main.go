package main

import (
	"fmt"
	"time"
)

func greet(greeting string) {
	fmt.Println(greeting)
}

/*
Creating a channel
make (chan <type of data>)

ex. make (chan int), make (chan float64), make (chan string), make (chan struct), etc

we can pass the created channels to the function that we intend to run as a goroutine
*/

func main() {
	doneChannel := make(chan bool)
	go slowGreetWithChannel("How... are... you...?", doneChannel)
	// here now we can read from the channel
	<-doneChannel // here we're now basically waiting for the channel to emit data

	// we can let the data flow into the void as above or
	// we could also print this data
	/*
		fmt.Println(<-doneChannel)
	*/
}

func slowGreetWithChannel(greeting string, doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(greeting)
	// here inside the function, once we're done with our operation
	// we can use the channel to send data through that channel to the place
	// where we started the goroutine
	doneChannel <- true
	// special arrow operator, NOTE: arrow always indicates the direction of data transfer
	// here data is being transfered into the channel.
}

// The above case demonstrated how a channel can be used to wait for the completion of a goroutine
// We can also use a channel to wait for the completion of multiple goroutines
// Go To: dealing_with_multiple_channels
