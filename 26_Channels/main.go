package main

import (
	"fmt"
	"sync"
)

func main() {

	myChannel := make(chan int, 1) // we have to specify the type of data that will be passed through the channel

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myChannel) //gives error i.e., DEADLOCK! - All goroutines are asleep - deadlock! AGAIN! (we thought we are listening here) Hence, we use goroutines.
	// myChannel <- 5
	// fmt.Println(<-myChannel) //gives error i.e., DEADLOCK! - All goroutines are asleep - deadlock!
	// This is because the channel is only allowing you to pass the value if somebody is listening to me. So, we cannot pass the value to the channel without somebody listening to it (fmt.Println(<-myChannel) in this case).

	wg.Add(2)
	//RECEIVE ONLY CHANNEL (if we mistakely write close having specified it as receive only channel we will be able to know that we made a mistake)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val,isChannelOpen := <-myChannel

		fmt.Println("Channel is open:", isChannelOpen)
		fmt.Println("Value from channel:", val)

		fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	//SEND ONLY CHANNEL
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// close(myChannel) // This will close the channel and will give error i.e., panic: send on closed channel
		myChannel <- 42
		myChannel <- 27 // This will give error IF the channel is unbuffered and can only hold one value at a time. Hence, it will give error i.e., fatal error: all goroutines are asleep - deadlock!
		myChannel <- 0 // It's tricky to tell on output if it is our zero or the zero from the channel being empty.
		close(myChannel) // works as all work is done and we are closing the channel
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
