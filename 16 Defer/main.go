package main

import "fmt"

func main() {
	defer myDefer()	
	defer fmt.Println("This will be printed at the end")
	defer fmt.Println("This will be printed second") //printed in reverse order they are deferred

	fmt.Println("SOme text")
}

func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}