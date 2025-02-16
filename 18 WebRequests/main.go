package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {

	res,err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", res) // *http.Response


	defer res.Body.Close() // This is important to close the connection, It's our responsibility to close the connection.

	dataBytes,err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Data type: %T\n", dataBytes) // []uint8
	fmt.Printf("Data type: %T\n", string(dataBytes)) // string


	fmt.Printf("%s\n", string(dataBytes))

}