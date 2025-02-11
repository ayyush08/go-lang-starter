package main

import "fmt"

func main() {
	// var ptr *int

	// fmt.Println("Value of ptr is: ", ptr) // Value of ptr is:  <nil>

	myNum := 34

	var ptr = &myNum // reference to myNum

	fmt.Println("Value of actual pointer is ",ptr)// Value of actual pointer is 0xc00000a0d8
	fmt.Println("Value of actual pointer is ",*ptr)// Value of actual pointer is  34

	*ptr *= 10 // multiply the value of myNum by 10

	fmt.Println("New value of myNum is ",myNum) // New value of myNum is  340
}