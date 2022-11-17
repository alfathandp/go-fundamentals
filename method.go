package main

import "fmt"

type Employee struct {
	firstName, lastName string
}

func (newEmployee Employee) fullName() string {
	return newEmployee.firstName + " " + newEmployee.lastName
}

func main()  {
	e := Employee {
		firstName: "alfathan",
		lastName: "dandy",
	}		
	fmt.Println(e.fullName())
}