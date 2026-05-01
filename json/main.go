package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are Learning JSON")
	person := Person{Name: "John", Age: 23, IsAdult: true}
	fmt.Println("Person Data is : ", person)

	// Convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling ", err)
		return
	}
	fmt.Println("Person data is : ", string(jsonData))
}
