package main

import "fmt"

func main() {
	fmt.Println("We are Learning Array in Golang")

	// 0 1 2 3 4
	// var name [5]string
	// name[0] = "Alex"
	// name[2] = "Tovino"

	// fmt.Println("Names of  Person is :", name)

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Number is :", numbers)

	// fmt.Println("Length of Numbers array is :", len(numbers))

	// fmt.Println("Value of name at 2 index is :", name[2])

	var name [5]string // by default [   ]
	name[2] = "Tovino"
	name[0] = "Alex"
	// var price [5]int //  by default 0
	fmt.Println("Name is :", name)
	fmt.Printf("Name Array is %q\n", name)

}
