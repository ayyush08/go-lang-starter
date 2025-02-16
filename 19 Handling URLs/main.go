package main

import (
	"fmt"
	"net/url"
)

const dummyUrl string = "https://example.com/api?param1=value1&param2=value2"

func main() {
	fmt.Println("URL:", dummyUrl)


	// Parse the URL

	result,err := url.Parse(dummyUrl)

	if err != nil {
		panic(err)
	}


	fmt.Println(result.Scheme) // https - The scheme is the protocol used to access the resource
	fmt.Println(result.Host) // example.com - The host is the domain name
	fmt.Println(result.Path) // /api - The path is the location of the resource on the server
	fmt.Println(result.RawQuery) // param1=value1&param2=value2 - The query string is the parameters passed to the resource
	fmt.Println(result.Port()) // "" - The port is the port number used to access the resource (provided after : in the host)

	queryParams:= result.Query() // store the query parameters in a map

	fmt.Printf("Type of queryParams: %T\n",queryParams) //url.Values

	fmt.Println(queryParams["param1"]) // [value1] - Access the value of a parameter by its key
	fmt.Println(queryParams["param2"]) // [value2]

	//we can also loop through the query parameters

	for _,val := range queryParams{
		fmt.Println(val)
	}
 
	partsOfUrl := &url.URL{ //remember that URL is a struct and must be passed as reference
		Scheme: "https",
		Host: "example.com",
		Path: "/api",
		RawQuery: "param1=value1&param2=value2",
	}

	anotherURL := partsOfUrl.String() //convert the URL struct back to a string

	fmt.Println(anotherURL) //https://example.com/api?param1=value1&param2=value2

}