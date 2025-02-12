package main

import "fmt"

func main() {

	//Remember: There's no inheritance in Go, No super or parent. GO is a simple language with simple constructs. It says what it does and does what it says.

	ayush := User{"Ayush", "ayush@gmail.com", true, 21}

	fmt.Println(ayush)
	fmt.Printf("User Ayush details are: %+v\n", ayush) //%+v will print the field names as well

	fmt.Printf("Name is %v and Email is %v.\n", ayush.Name, ayush.Email)

}

type User struct { //"User" remember to capitalize the first letter
	Name   string
	Email  string
	Status bool
	Age    int
}
