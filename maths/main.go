package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {

	// var a int = 10
	// var b float64 = 9
	// var c float64 = 9.4

	// fmt.Println("Addition: ", a+b) // error as a and b are of different types
	// fmt.Println("Addition: ", a+int(b)) // type conversion
	// fmt.Println("Addition: ", a+int(c)) // type conversion

	
	//Generating Random Numbers
	// rand.New(rand.NewSource(time.Now().UnixNano())) - always gives the same random number as the seed is same
	// fmt.Println(rand.Intn(5)) // generates a random number between 0 and 4 when used without seeding

	//random number generation using crypto/rand

	randomVar, _ := rand.Int(rand.Reader,big.NewInt(5)) // generates a random number between 0 and 4
	fmt.Println(randomVar)
	//with crypto the randomness is much accurate than math/rand
}
