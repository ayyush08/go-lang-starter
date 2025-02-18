package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"` //we can use tags to change the key name in JSON data
	Price    int      `json:"coursePrice"`
	Platform string   `json:"coursePlatform"`
	Password string   `json:"-"`                    //we can use "-" to hide the key in JSON data
	Tags     []string `json:"courseTags,omitempty"` //we can use omitempty to hide the key if the value is nil, BUT BE CAREFUL ABOUT THE SPACE AFTER COMMA, THERE IS NONE
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	courses := []course{
		{"ReactJS", 299, "Udemy", "password", []string{"web-dev", "frontend"}},
		{"Django", 199, "Udemy", "password", []string{"web-dev", "backend"}},
		{"Flask", 199, "Udemy", "password", nil},
	}

	//package this data as JSON data

	finalJson, err := json.Marshal(courses) // Marshal implements the encoding/json package to convert the data to JSON format

	finalJsonWithIndent, err2 := json.MarshalIndent(courses, "", "\t") // MarshalIndent is used to make the JSON data more readable by adding indentation, the second argument is the prefix meaning the character to be added before each line and the third argument is the indentation character meaning the character to be added before each level of indentation.

	if err2 != nil {
		panic(err2)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON data: %s\n", finalJson) //hard to read(so we use MarshalIndent), nil converts to null to make it more readable for consumers

	fmt.Printf("JSON data with indentation: %s\n", finalJsonWithIndent) //We notice that the JSON data is more readable now but we see that the keys start with a capital letter which is not a good practice in JSON data or maybe inside API "Name" is needed as "course_name" so we use tags to change the key name in JSON data

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
                "courseName": "ReactJS",
                "coursePrice": 299,
                "coursePlatform": "Udemy",
                "courseTags": [
                        "web-dev","frontend"
                ]
        }
		`)

	var course course

	//check is json is valid

	isValidJson := json.Valid(jsonDataFromWeb)

	if !isValidJson {
		fmt.Println("Invalid JSON data")
		return
	}

	fmt.Println("JSON data is valid")
	json.Unmarshal(jsonDataFromWeb,&course) // Unmarshal is used to convert JSON data to Go data . & because sometimes when we pass a struct to a function, it is passed as a copy and not as a reference so we use & to pass it as a reference

	fmt.Printf("Course: %#v\n", course) // %#v is used to print the struct with field names

	//some cases where you just want to add data to key value pairs and not to a struct

	var myOnlineData map[string]interface{} //An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods. So here we are defining a map with string keys and interface values

	json.Unmarshal(jsonDataFromWeb,&myOnlineData)

	fmt.Printf("Online data: %#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("Key: %s, Value: %v and Type: %T\n",k,v,v)
	}
}
