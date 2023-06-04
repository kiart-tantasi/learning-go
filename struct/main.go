package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) PrintName() {
	fmt.Println("Hello ! my name is", p.FirstName, p.LastName)
}

type SomeFields struct {
	a, b string
	c    int
}

func (s SomeFields) UpdateBWithoutPointer(newB string) {
	s.b = newB
}
func (s *SomeFields) UpdateB(newB string) {
	s.b = newB
}

func main() {
	// normal struct
	person := Person{
		FirstName: "Hello",
		LastName:  "World",
	}
	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
	fmt.Println("")

	// use struct's method
	person.PrintName()
	fmt.Println("")

	// struct does not require values for all fields
	someFields := SomeFields{
		a: "a field",
	}
	fmt.Println("field that has value:", someFields.a)
	fmt.Println("field that does not have value(string):", someFields.b)
	fmt.Println("field that does not have value(int):", someFields.c)
	fmt.Println("")

	// pointer
	personPointer := &person
	fmt.Println("from pointer")
	fmt.Println(personPointer.FirstName)
	fmt.Println("")

	// update same value without / with pointer
	// person.UpdateFirstNameButFail("new name!")
	someFields.UpdateBWithoutPointer("new value !")
	fmt.Println("update without pointer:", someFields.b)
	someFields.UpdateB("new value !")
	fmt.Println("update with pointer:", someFields.b)
}
