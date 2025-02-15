package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)

	//Type 1

	for d:=0; d<len(days); d++ { //no such thing as ++d
		// fmt.Println(days[d])
	}

	//Type 2

	// for i := range days{ // i here returns the index of the slice
		// fmt.Println(days[i])
	// }

	//Type 3

	// for index, day := range days{
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }


	rougeValue := 1

	for rougeValue < 10{ //kind of like a while loop
		if rougeValue==3{
			goto lco
		}

		if(rougeValue==5){
			rougeValue++
			continue // only continue does not skip as in other languages we need to increment the value
		}

		fmt.Println(rougeValue)
		rougeValue++
	}

	lco:
		fmt.Println("Jumping to LCO") //goto label
}