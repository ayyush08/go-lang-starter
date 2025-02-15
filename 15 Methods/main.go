package main

import "fmt"

//In Go, since we don't have classes, we have structs. We can attach methods to structs.

func main() {
	ayush := User{"Ayush","ayush@gmail.com",false, 25}
	fmt.Println("Email before change ", ayush.Email)
	ayush.GetStatus()
	ayush.NewMail()
	fmt.Println("Email after change ", ayush.Email) //Email is still the same because we're not changing the original struct, we're changing a copy of it but we can change it by using pointers
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){ //g should be uppercase if you're planning to  export it
	fmt.Println("Is user active? ", u.Status)
}

func (u User) NewMail(){
	u.Email = "cura@ga.com"
	fmt.Println("New email is ", u.Email)
}