package main

import (
	"fmt"
	"net/http"
	"sync"
)


var waitGroupVariable sync.WaitGroup //wait group in sync package is used to wait for go routines to come back and once they come it is your responsibility to call done() on waitGroupVariable to let it know that the go routine has come back

func main() {

	// go greeter("Hello") // fire up a thread which is responsible further for execution of greeter function but main function will never wait for it to come back and print hello (we can use sleep to wait for it to come back)
	//just prints There and exits
	// greeter("There")

	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
		"https://instagram.com",
		"https://x.com",
	}

	fmt.Println("Starting the go routines")
	for _, website := range websiteList {
		// getStatusCode(website) // this will take a lot of time to get the status code for all the websites as it is a blocking call and it will wait for the response to come back
		// we can use go routines to make it non blocking


		go getStatusCode(website) // this will fire up a go routine for each website and will not wait for the response to come back BUT just adding 'go' keyword will not make it non blocking as main function will exit before the go routines come back since we are not waiting for them to come back
		waitGroupVariable.Add(1) // increment the wait group counter to let it know that we are waiting for one more go routine to come back
	}
	waitGroupVariable.Wait() // wait for all the go routines to come back - always written at end of any method otherwise it will block the main thread
	fmt.Println("All go routines have come back")	
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond) // not an ideal way to wait for a go routine to come back but just for the sake of understanding. Usually we use 'sync' package to wait for go routines to come back
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitGroupVariable.Done() // this will be called when the go routine comes back and will decrement the wait group counter
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Endpoint is not reachable: ", err)
		return
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
