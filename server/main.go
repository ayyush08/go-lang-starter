package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
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


func PerformPostRequest() {
	const myurl = "http://localhost:5000/post"

	//fake json payload

	reqBody := strings.NewReader(`
	{
	"name":"Ichigo Kurosaki",
	"age":20,
	"job":"Soul Reaper"
	}`) // we can create any kind of data inside these backticks

	res,err := http.Post(myurl, "application/json",reqBody)

	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.StatusCode)
	fmt.Println("Response content length:", res.ContentLength)

	content, _ := io.ReadAll(res.Body)

	fmt.Println("Response body:", string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:5000/postform"


	//fake form data

	formData := url.Values{} // as value comes in, we can add it to the form data

	formData.Add("name","Monkey D. Luffy")
	formData.Add("age","19")
	formData.Add("job","Pirate King")
	formData.Add("crew","Straw Hat Pirates")
	

	res,err := http.PostForm(myurl,formData) // here we don't need to specify the content type as it is already specified in the function

	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println("Response body:", string(content))

}