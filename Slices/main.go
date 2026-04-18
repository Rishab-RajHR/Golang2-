package main

import "fmt"

func main() {

	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10, 22, 55, 89, 34)
	// fmt.Println("Number :", numbers)
	// fmt.Printf("Number has data type : %T\n", numbers)
	// fmt.Println("Length : ", len(numbers))

	// name := []string{}
	// fmt.Println("name :", name)

	numbers := make([]int, 3, 5)

	// When we increase the numbers more than the capacity then capacity will be doubled

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 10)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	stock := make([]int, 0)
	fmt.Println("Slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))
}
