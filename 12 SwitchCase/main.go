package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	diceNumber := rand.Intn(6) + 1


	switch diceNumber{
		case 1:
			println("Dice number is 1")
			//no need to write break statement, butfor fallthrough, we need to write it
			fallthrough
		case 2:
			println("Dice number is 2")
		case 3:
			println("Dice number is 3")
		case 4:
			println("Dice number is 4")
		case 5:
			println("Dice number is 5")
		case 6:
			println("Dice number is 6")
		default:
			println(diceNumber, "is invalid")
	}

}