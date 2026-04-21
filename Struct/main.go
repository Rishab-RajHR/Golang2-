package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var alex Person
	fmt.Println("Alex Person : ", alex)
	alex.FirstName = "Alex"
	alex.LastName = "Pandian"
	alex.Age = 24
	fmt.Println("Alex Person : ", alex)
}
