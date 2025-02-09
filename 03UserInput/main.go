package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the age: ")

	//Comma ok syntax aka Err ok syntax - we don't have try catch in go. Go says treat error as a value.

	age, _ := reader.ReadString('\n')
	fmt.Println("Entered name is: ", age)
	fmt.Printf("Type of name is: %T\n", age)

}
