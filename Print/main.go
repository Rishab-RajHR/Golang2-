package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.823443

	// Println adds a new line
	fmt.Println("age: ", age, " height: ", height, " name: ", name)
	fmt.Println("Hello World")

	// Printf => format specifier
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.4f\n", height)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name is %s\n", name)

	// Multiple specifier in one line
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
