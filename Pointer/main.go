package main

import "fmt"

func main() {
	var num int
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println("Num has value: ", num)
	fmt.Println("Pointer contains: ", ptr)
}
