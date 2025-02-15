package main

import "fmt"

func main() {
	fmt.Println("SOme text")
	greeter()

	result := adder(3, 5)

	fmt.Println(result)

	result2, statement := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(result2,statement)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values{
		total += val
	}
	return total , " is the total"
}


func greeter(){
	fmt.Println("Hello Curator")
}