package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayyush08/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":3000",r))

	fmt.Println("Server is running on port 3000")

}
