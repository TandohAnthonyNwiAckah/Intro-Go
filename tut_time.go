package main

import (
	"fmt"
	"time"
)

func TutTime() {

	fmt.Println("Time in Go:")

	timeNow := time.Now()
	fmt.Println("Current Time:", timeNow)
	fmt.Println("Current Time:", timeNow.Format("01-02-2006"))
	fmt.Println("Current Time:", timeNow.Format("01-02-2006 Monday"))
	fmt.Println("Current Time:", timeNow.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.March, 01, 10, 10, 0, 0, time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday "))
}
