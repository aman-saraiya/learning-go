package main

import (
	"fmt"
	"time"
)

func greetWithChannel(greeting string, doneChannel chan bool) {
	fmt.Println(greeting)
	doneChannel <- false
}

func slowGreetWithChannel(greeting string, doneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulating a slow time taking task
	fmt.Println(greeting)
	doneChannel <- true
}

/***************************************************************/
func usingSameChannelMultipleFunctions() {
	doneChannel := make(chan bool)
	go greetWithChannel("Nice to meet you", doneChannel)
	go greetWithChannel("How are you?", doneChannel)
	go slowGreetWithChannel("How ... are ... you ...?", doneChannel)
	go greetWithChannel("I hope you're having a great day", doneChannel)
	fmt.Println(<-doneChannel) // will get the value from the function that finishes first, (could be any of above)
	fmt.Println(<-doneChannel) // will get the value from the function that finishes second, (could be any of remaining)
	fmt.Println(<-doneChannel) // and so on
	fmt.Println(<-doneChannel)
}

// using multiple lines like above <-doneChannel is not very scalable and we always have to update the
// number of reads from the channel whenever we add another go routine, since that is not ideal there's
// an alternative way of handling this using slice.

/***************************************************************/
func usingDifferentChannelsMultipleFunctions() {
	doneChannels := make([]chan bool, 4)
	doneChannels[0] = make(chan bool)
	doneChannels[1] = make(chan bool)
	doneChannels[2] = make(chan bool)
	doneChannels[3] = make(chan bool)

	go greetWithChannel("Nice to meet you", doneChannels[0])
	go greetWithChannel("How are you?", doneChannels[1])
	go slowGreetWithChannel("How ... are ... you ...?", doneChannels[2])
	go greetWithChannel("I hope you're having a great day", doneChannels[3])

	for _, doneChannel := range doneChannels {
		<-doneChannel
	}
}

/***************************************************************/
// Again managing such slice of channels can be cumbersome, we can also
// use another feature
func alternativeWaySingleChannelMultipleFunctions() {
	doneChannel := make(chan bool)
	go greetWithChannel("Nice to meet you", doneChannel)
	go greetWithChannel("How are you?", doneChannel)
	go slowGreetWithChannel("How ... are ... you ...?", doneChannel)
	go greetWithChannel("I hope you're having a great day", doneChannel)

	// we can directly use for loop on a channel

	// this is a feature that is supported by go
	for channel := range doneChannel {
		// here inside the loop we will now not get the channels but he values emitted
		// by that channel
		fmt.Println(channel)
	}
}

/* Here we will get the following error as go doesn't know when this channel is out of values
we are simply looping through all the values and are waiting for new values to be emitted but
eventually there will be no values left and hence we get this error

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
		D:/Shift/learning-go/main.go:136 +0x185
exit status 2
*/

/***************************************************************/
/*
We could get rid of this error by explicitly closing the channel once we're done.
But this approach would only work if we know what is the last value or what function
emits the last value into the channel.

In the current example we know it's the slowGreet function so we could close the channel in it
as close(doneChannel). We cannot emit or read any data from the channel once closed
*/

func slowGreetWithChannelClose(greeting string, doneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulating a slow time taking task
	fmt.Println(greeting)
	doneChannel <- true
	close(doneChannel)
}

func workingAlternativeWaySingleChannelMultipleFunctions() {
	doneChannel := make(chan bool)
	go greetWithChannel("Nice to meet you", doneChannel)
	go greetWithChannel("How are you?", doneChannel)
	go slowGreetWithChannelClose("How ... are ... you ...?", doneChannel)
	go greetWithChannel("I hope you're having a great day", doneChannel)

	// we can directly use for loop on a channel

	// this is a feature that is supported by go
	for channel := range doneChannel {
		// here inside the loop we will now not get the channels but he values emitted
		// by that channel
		fmt.Println(channel)
	}
}

/***************************************************************/

func main() {
	// using same channel for multiple functions
	/*
		usingSameChannelMultipleFunctions()
	*/

	// using different channels for different functions
	/*
		usingDifferentChannelsMultipleFunctions()
	*/
	// here the channel outputs are in order, the for loop won't move further until
	// doneChannels[0] is processed, then doneChannels[1], and so on.
	// even if function with doneChannels[1] is processed first it will pass
	// the value to the channel and finish but the value will be taken out
	// only after the loop reaches till doneChannels[1]

	/*
		alternativeWaySingleChannelMultipleFunctions()
	*/

	workingAlternativeWaySingleChannelMultipleFunctions()

}
