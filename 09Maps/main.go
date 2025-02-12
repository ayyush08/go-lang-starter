package main

import "fmt"

func main() {

	languages := make(map[string]string) //map[keyType]valueType

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "Typescript"

	fmt.Println("List of all languages:", languages) 
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB") //deleting an element from map (can also be used for slices)
	fmt.Println("List of all languages after deleting Ruby:", languages)

	//interesting thins about loops in golang

	for key,value := range languages{ // walrus operator is responsible for all comma ok syntax 
		fmt.Printf("For key %v, value is %v\n", key, value) //%v is used to print the value of the variable
	}
	for _,value := range languages{ // walrus operator is responsible for all comma ok syntax 
		fmt.Printf("value is %v\n", value) 
	}

}