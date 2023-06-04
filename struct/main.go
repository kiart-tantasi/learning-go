package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

func (p Person) PrintName() {
	fmt.Println("Hello ! my name is", p.firstName, p.lastName)
}

type SomeFields struct {
	a, b string
	c int
}
func (s SomeFields) UpdateBWithoutPointer(newB string) {
	s.b = newB
}
func (s *SomeFields) UpdateB(newB string) {
	s.b = newB
}

func main() {
	// normal struct
	person := Person {
		firstName: "Hello",
		lastName: "World",
	}
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	fmt.Println("")

	// use struct's method
	person.PrintName()
	fmt.Println("")

	// struct does not require values for all fields
	someFields := SomeFields {
		a: "a field",
	}
	fmt.Println("field that has value:", someFields.a)
	fmt.Println("field that does not have value(string):", someFields.b)
	fmt.Println("field that does not have value(int):", someFields.c)
	fmt.Println("")

	// pointer
	personPointer := &person
	fmt.Println("from pointer")
	fmt.Println(personPointer.firstName)
	fmt.Println("")

	// update same value without / with pointer
	// person.UpdateFirstNameButFail("new name!")
	someFields.UpdateBWithoutPointer("new value !")
	fmt.Println("update without pointer:", someFields.b)
	someFields.UpdateB("new value !")
	fmt.Println("update with pointer:", someFields.b)
}
