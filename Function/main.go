package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a, b int) int {
	return a + b
}

// func add(a, b int) (result int) {
// 	  result = a + b
// 		return
// }

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("We are Learning Function in Golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("Add of two number is :", ans)

	data := multiply(4, 5)
	fmt.Println("Multiplication of two numbers is :", data)
}
