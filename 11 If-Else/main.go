package main

import "fmt"

func main() {

	loginCount := 10

	var result string
	if loginCount < 10 { // not allowed to move this curly brace to the next line
		result = "Regular User"
	} else if loginCount >10 {
		result = "Super User"
	} else {
		result = "No User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num:= 3; num< 5 {
		fmt.Println(num, "is less than 5")
	} else{
		fmt.Println(num, "is not less than 5")
	}

}