package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p Person) PrintName() {
	fmt.Println("Hello ! my name is", p.firstName, p.lastName)
}

type NestedStruct struct {
	value *string
}
type SomeStruct struct {
	a, b         string
	c            int
	d            *string
	nestedStruct NestedStruct
}

func (s SomeStruct) UpdateBWithoutPointer(value string) {
	s.b = value
}
func (s *SomeStruct) UpdateB(value string) {
	s.b = value
}

func main() {
	// normal struct
	person := Person{
		firstName: "Hello",
		lastName:  "World",
	}
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	fmt.Println("")

	fmt.Println("==============================================")

	// use struct's method
	person.PrintName()
	fmt.Println("")

	fmt.Println("==============================================")

	// struct does not require values for all fields
	nestedStructValue := "OLD VALUE"
	nestedStruct := NestedStruct{value: &nestedStructValue}
	foo := SomeStruct{
		a:            "a field",
		nestedStruct: nestedStruct,
	}
	fmt.Println("field that has value:", foo.a)
	fmt.Println("field that does not have value(string):", foo.b)
	fmt.Println("field that does not have value(int):", foo.c)

	fmt.Println("==============================================")

	// pointer
	personPointer := &person
	fmt.Println("from pointer")
	fmt.Println(personPointer.firstName)

	// failed
	// person.UpdateFirstNameButFail("new name!")

	fmt.Println("==============================================")

	// update same value without pointer
	foo.UpdateBWithoutPointer("new value ! (without pointer)")
	fmt.Println(foo.b)

	// update same value with pointer
	foo.UpdateB("new value ! (with pointer)")
	fmt.Println(foo.b)

	fmt.Println("==============================================")

	// use external function to modify a struct directly
	updateStruct(foo, "new value ! (struct)")
	pointer := foo.d
	if pointer != nil {
		fmt.Println(*pointer)
	} else {
		fmt.Println("nil (struct)")
	}

	// modify a struct pointer
	updateStructPointer(&foo, "new value ! (struct pointer)")
	pointer = foo.d
	if pointer != nil {
		fmt.Println(*pointer)
	} else {
		fmt.Println("nil (struct pointer)")
	}

	// modify nested struct
	// NOTE: This is exceptional case where you can modify a nested struct inside struct without passing struct as pointer to the function
	updateNestedStruct(foo, "new value ! (nested struct)")
	pointer = foo.nestedStruct.value
	if pointer != nil {
		fmt.Println(*pointer)
	} else {
		fmt.Println("nil (nested struct)")
	}

	fmt.Println("==============================================")
}

func updateStruct(s SomeStruct, value string) {
	if s.d != nil {
		*s.d = value
	}
	newPointer := &value
	s.d = newPointer
}

func updateStructPointer(s *SomeStruct, value string) {
	if s.d != nil {
		*s.d = value
	}
	newPointer := &value
	s.d = newPointer
}

func updateNestedStruct(s SomeStruct, value string) {
	if s.nestedStruct.value == nil {
		fmt.Println("...I did try to modify nested sturct but it does not work...")
		s.nestedStruct.value = &value
		return
	}
	*s.nestedStruct.value = value
}
