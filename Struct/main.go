package main

import "fmt"

// int string bool Person Contact Address
// var name Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var alex Person
	// fmt.Println("Alex Person : ", alex)
	alex.FirstName = "Alex"
	alex.LastName = "Pandian"
	alex.Age = 24
	// fmt.Println("Alex Person : ", alex)

	// 2nd Method
	person1 := Person{
		FirstName: "Tovino",
		LastName:  "Thomas",
		Age:       26,
	}
	fmt.Println(" Person 1 : ", person1)

	// new keyword
	var person2 = new(Person)
	person2.FirstName = "Basil"
	person2.LastName = "Joseph"
	person2.Age = 26

	// fmt.Println("Person 2 : ", person2.FirstName)

	// fmt.Println("Age of Alex is ", alex.Age)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Alex",
		LastName:  "Pandian",
		Age:       28,
	}
	employee1.Person_Contact.Email = "alex12@gmail.com"
	employee1.Person_Contact.Phone = "978875434"

	employee1.Person_Address = Address{
		House: 12,
		Area:  "Tanjavur",
		State: "Jharkhand",
	}

	// fmt.Println("Employee 1 : ", employee1)

	fmt.Println("Employee 1 : ", employee1.Person_Contact.Email)
}
