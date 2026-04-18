package main

import "fmt"

func main() {
	fmt.Println("Learn Golang")

	var name string = "Alex"
	var version = 76
	fmt.Println(name)
	fmt.Println(version)

	var money int = 67000
	var currency = 3489
	fmt.Println(money)
	fmt.Println("currency: ", currency)

	var dimension float64 = 87.62
	fmt.Println(dimension)

	var decided bool = false
	decided = true
	fmt.Println(decided)

	var person = "Alex Pandian"
	fmt.Println(person)

	const pi = 67.12
	fmt.Println(pi) // cannot change afterwrds

	// Printing variable through other method
	person1 := "Alex Joseph"
	fmt.Println(person1)

	// Capitalization
	// Variable declared with capital can be accessed outside the file
	// Variable declared without capital cannot be accessed outside the file

	var Public = "data is important"
	var private = "data is private"

	fmt.Println(Public)
	fmt.Println(private)

}
