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

	// Check key present or not
	grades, exists := studentGrades["David"]
	fmt.Println("Grades of David : ", grades)
	fmt.Println("Exists of David : ", exists)

	//  fmt.Println("Marks of David : ", studentGrades["David"])

	// Check key present or not
	Grades, Exists := studentGrades["Alex"]
	fmt.Println("Grades of Alex : ", Grades)
	fmt.Println("Exists of Alex : ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	// Map declare and initialize at the same time
	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}

	for index, value := range person {
		fmt.Printf("key is %s and marks is %d\n", index, value)
	}
}
