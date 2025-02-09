package main

import (
	"fmt"
)

// jwtToken := 5000 // This will give an error as we can't use walrus operator outside of any method, instead use var or const

const JwtToken string = "dfafaf" // This will work as const can be defined outside of any method
//Also, when first letter of variable is capital, it means it is public and can be accessed outside of package

func main() {
	var username string = "Ayush"
	fmt.Println("Hello", username)
	fmt.Printf("Variable is of type : %T\n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T\n", isLoggedIn)
	var smallValue uint8 = 5
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T\n", smallValue)
	var smallFloat float64 = 87.123456789
	fmt.Println("Hello", smallFloat)
	fmt.Printf("Variable is of type : %T\n", smallFloat)

	// Default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T\n", anotherVariable)

	//implicit type

	var website = "golang.org"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T\n", website)
	// website = 5 // This will give an error as lexer have already defined website as string

	//No var style

	numberOfUsers := 300000 // walrus operator
	fmt.Println("Number of users", numberOfUsers)

	fmt.Println(JwtToken)
	fmt.Printf("Type of JwtToken %T", JwtToken)
}
