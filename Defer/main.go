package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the Program")
	data := add(5, 6)
	fmt.Println("Data is : ", data)
	fmt.Println("Middle of the Program")
	fmt.Println("End of the Program")
}
