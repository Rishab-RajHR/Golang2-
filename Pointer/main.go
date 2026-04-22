package main

import "fmt"

func main() {
	// var num int
	num := 2

	// var ptr *int
	// ptr = &num
	ptr := &num

	// fmt.Println("Num has value: ", num)
	fmt.Println("Pointer contains: ", ptr)
	fmt.Println("Data contains through Pointer: ", *ptr) // 2

	var pointer *int // default pointer == nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}
}
