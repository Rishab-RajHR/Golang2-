package main

import "fmt"

func main() {
	var num int = 42
	fmt.Println("Number is ", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Println("Data is ", data)
	fmt.Printf("Type of Data is %T\n", data)
}
