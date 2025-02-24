package main

import (
	"fmt"
	"time"
)

func main() {

	go greeter("Hello") // fire up a thread which is responsible further for execution of greeter function but main function will never wait for it to come back and print hello (we can use sleep to wait for it to come back)
	//just prints There and exits
	greeter("There")

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3*time.Millisecond) // not an ideal way to wait for a go routine to come back but just for the sake of understanding. Usually we use 'sync' package to wait for go routines to come back
		fmt.Println(s)
	}
}
