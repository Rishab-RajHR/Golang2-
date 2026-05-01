package main

import "fmt"

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are Learning JSON")
	person := Person{Name: "John", Age: 23, IsAdult: true}
	fmt.Println("Person Data is : ", person)
}
