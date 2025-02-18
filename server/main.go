package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
	
}

func PerformGetRequest() {
	const myurl = "http://localhost:5000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	//don't forget to close the response body, IT IS YOUR JOB
	defer res.Body.Close()

	fmt.Println("Response status:", res.StatusCode)
	fmt.Println("Response content length:", res.ContentLength)

	//read the response body
	
	
	//we can also use strings.Builder to build the response body
	
	content, _ := io.ReadAll(res.Body)

	// fmt.Println("Response body:", string(content)) //string here is a data type
	
	var responseString strings.Builder

	byteCount ,_ := responseString.Write(content)

	fmt.Println("Bytes written to responseString:", byteCount)
	fmt.Println(responseString.String()) //string here is a method, it converts the raw or byteform data to a string. Earlier string worked and both are good but this is more efficient and recommended as it comes with more features


}
