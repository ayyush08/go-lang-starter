package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Model for course - file (supposed to go in a file)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //Pointer to Author struct
}

type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

//fakeDB

var courses []Course

//middleware or helpers - file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Cura API</h1>"))

}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //whatever we write inside Encode will be treated as json and will be thrown back to the person requesting
	
}
