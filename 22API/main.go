package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Println("Api starts")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 100, Author: &Author{Fullname: "John Doe", Website: "www.johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Angular", CoursePrice: 200, Author: &Author{Fullname: "Jane Doe", Website: "www.janedoe.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Vue", CoursePrice: 300, Author: &Author{Fullname: "Michael Doe", Website: "www.michaeldoe.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":3000", r))
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

	//check duplicate course name

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			fmt.Println("Course already exists")
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the url or request

	params := mux.Vars(r) //grabbing the parameters from the request

	fmt.Println("Params: ", params)

	//loop through the courses and find the course with the id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //... for variadic operation

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			fmt.Println("Course: ", course)
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send response if course not found
	json.NewEncoder(w).Encode("Course not found with the id " + params["id"])
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the url or request

	params := mux.Vars(r) //grabbing the parameters from the request

	fmt.Println("Params: ", params)

	if params["id"] == "" {
		json.NewEncoder(w).Encode("Please provide the id")
		return
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the id " + params["id"])
}
