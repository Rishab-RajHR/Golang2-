package main

import "fmt"

// func divide(a, b int) int {
// 	return a / b
// }

// func divide(a, b float64) float64 {
// 	return a / b
// }

// Error Handling in Function
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator must not be zero")
// 	}
// 	return a / b, nil
// }

// Error Handling in String
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator must not be zero"
	}
	return a / b, "nil"
}

func main() {
	fmt.Println("Started Error Handling in the Functions")
	// ans := divide(10, 4)
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error Handling")
	// }
	ans, _ := divide(10, 0)
	fmt.Println("Division of two numbers is ", ans)
}
