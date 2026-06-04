package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{
		"Gus",
		"Gomez",
		1,
	}

	emp2 := Employee{
		firstName: "John",
		lastName:  "Jameson",
		id:        2,
	}

	var emp3 Employee
	emp3.firstName = "Jone"
	emp3.lastName = "Jane"
	emp3.id = 3

	fmt.Println(emp1, emp2, emp3)
}
