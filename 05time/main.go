package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println("Present Time: ", presentTime)

	fmt.Println("Present Time formatted: ", presentTime.Format("01-02-2006 15:04")) //strict rule for date and time formatting

	createdDate := time.Date(2025, time.January,1 ,02,12,12,12, time.Local).Format("01-02-2006 Monday")

	fmt.Println("Created Date: ", createdDate)

}
