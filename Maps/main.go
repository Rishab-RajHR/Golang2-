package main

import "fmt"

func main() {

	// name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Alex"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of Bob : ", studentGrades["Bob"])

	// Change the marks
	studentGrades["Bob"] = 100
	fmt.Println("New marks of Bob : ", studentGrades["Bob"])

	//  Delete Bob Marks
	delete(studentGrades, "Bob")
	fmt.Println("Marks of Bob : ", studentGrades["Bob"])
}
