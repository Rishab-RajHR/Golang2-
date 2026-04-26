package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	formatted := currentTime.Format("2006-02-04, 3:04 PM")
	fmt.Println("Formatted time: ", formatted)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time:", formatted_time)
}
