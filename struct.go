package main

import "fmt"


type person struct {
	firstName string
	lastName string
	Age int
}

func main()  {
	// long declaration 
	var person0 person
	person0.firstName = "alfathan"
	person0.lastName = "person 0"
	person0.Age = 21
	fmt.Println(person0.firstName, person0.lastName, person0.Age)

	// long declaration dengan assigned value
	var person1 = person{"alfathan", "person 1", 21}
	fmt.Println(person1)
	
	// long declaration dengan assigned value tiap nama field
	var person2 = person {
		firstName: "alfathan",
		lastName: "person 2",
		Age: 21,
	} 
	fmt.Println(person2)

	//short declaration
	person3 := person{"alfathan", "person 3", 21}
	fmt.Println(person3)

	//short declaration with "new" keyword
	person4 := new(person)
	person4.firstName = "alfathan"
	person4.lastName = "person 4"
	person4.Age = 21
	fmt.Println(*person4)
}