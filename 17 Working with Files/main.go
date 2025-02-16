package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "This needs to be written to a file - from Go"

	file, err := os.Create("./myFile.txt")

	if err != nil {
		fmt.Println("Error creating file")
		panic(err) // panic is a built-in function that stops the ordinary flow of control and begins panicking.
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("Error writing to file")
		panic(err)
	}

	fmt.Printf("Wrote a string of length %d to file\n", length)

	defer file.Close() //defer is recommended to use here to close the file after the function completes
	readFile("./myFile.txt")

}


func readFile(filename string){
	// ioutil.ReadFile(filename) - deprecated

	dataBytes, err := os.ReadFile(filename)

	checkNilErr(err)

	dataString := string(dataBytes)

	fmt.Println("Data read from file:", dataString)
}


func checkNilErr(err error){ //common practice to check for nil errors
	if err!=nil{
		panic(err)
	}
}