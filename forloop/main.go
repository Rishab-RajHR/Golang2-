package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("Numbers is :", i)
	}

	// For infinite Loop stop using the break
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 2 {
			break
		}
	}

	// range keyword => index and value
	numbers := []int{11, 42, 83, 14, 75}
	for index, value := range numbers {
		fmt.Printf("Index of Numbers is %d, and value is %d\n", index, value)
	}

	// For accessing the String value
	data := "Hello, World!"
	for index, char := range data {
		fmt.Printf("Index of Data is %d, and character  is %c\n", index, char)
	}

}
