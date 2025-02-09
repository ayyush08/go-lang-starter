package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	// ReadString will read until the first newline character

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating our pizza: ", rating)

	// numRating = input+1 // This will throw an error because rating is a string and not an integer

	// numRating, err := strconv.ParseFloat(rating, 64) // error as rating is trailed with '\n'
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64) 

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
