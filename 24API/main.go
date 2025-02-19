package main



//Model for course - file (supposed to go in a file)

type Course struct {
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"` //Pointer to Author struct
}


type Author struct{
	Fullname string `json:"fullName"`
	Website string `json:"website"`
}

//fakeDB

var courses []Course

//middleware or helpers - file

func (c *Course) IsEmpty() bool{
	return c.CourseId=="" && c.CourseName ==""
}




func main() {
	
}