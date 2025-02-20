package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

//	func (c *Course) IsEmpty() bool {
//		return c.CourseId == "" && c.CourseName == ""
//	}
func (c *Course) IsEmpty() bool {
	return c.CourseId == ""
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//grab the id from the url or request

	params := mux.Vars(r) //grabbing the parameters from the request

	fmt.Println("Params: ", params)

	//loop through the courses and find the course with the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the id " + params["id"])
	return

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide data")
		return
	}

	// what about {} in the body

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No course data provided")
		return
	}

	//generate unique id, string
	//append the course to the courses

	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
