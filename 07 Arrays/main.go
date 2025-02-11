package main

import "fmt"

func main() {
	//In Go lang, there is not so much use of arrays
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[2] = "Orange" // when not assigned, it will be empty string and in output it looks like an extra empty space
	fruitList[3] = "Peach"

	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List Length:", len(fruitList)) //gives length of array defined(4) not the number of elements assigned(3)

	var vegList = [3]string{"Potato", "Tomato", "Brinjal"} //short hand declaration

	fmt.Println("Veg List:", vegList)
	fmt.Println("Veg List Length:", len(vegList))

}
