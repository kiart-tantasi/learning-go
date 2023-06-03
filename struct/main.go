package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

func main() {
	person := Person {
		firstName: "Hello",
		lastName: "World",
	}
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	
	// pointer
	personPointer := &person
	fmt.Println("from pointer")
	fmt.Println(personPointer.firstName)
}
