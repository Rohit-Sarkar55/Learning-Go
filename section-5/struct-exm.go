package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
	JoinedAt  time.Time
}

func newEmployee(id int, firstName, lastName, position string, salary float64, isActive bool) *Employee {
	return &Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}
func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	joe := newEmployee(2, "Joe", "Hart", "Manager", 3400, true)

	fmt.Printf("%+v\n", jane)

	fmt.Println(joe.ID)
	fmt.Println(joe.FirstName)
	fmt.Println(joe.LastName)

	fmt.Println(joe) // joe pointer

	fmt.Println(&joe)
}
