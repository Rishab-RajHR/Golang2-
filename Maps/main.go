package main

import "fmt"

func main() {

	// name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Alex"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks if Bob : ", studentGrades["Bob"])
}
