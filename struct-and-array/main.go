package main

import "fmt"

type Car struct {
	id   int32
	name string
}

func main() {
	// init a new car
	var cars []Car
	car1 := Car{
		id:   12345,
		name: "New Car",
	}
	cars = append(cars, car1)

	// pointer
	carPointer := &car1
	fmt.Println(carPointer)

	// dereference
	fmt.Println((*carPointer).name)

	// array is fixed-sized in go
	fmt.Println("[array]")
	array := [2]int{9, 10}
	fmt.Println("first element in array is", array[0])
	// array = append(array, 11) // this will give error because array cannot be used with append function
	fmt.Println("=========================")

	// slice (dynamically-sized array)
	nums := []int{1, 2, 3, 4, 5, 6}

	// length and capacity before adding
	fmt.Println("[slice]")
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d\n", cap(nums))
	fmt.Println("=========================")

	// appending
	fmt.Println("[appended]")
	nums = append(nums, 7)
	// length and capacity after adding
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d (*notice that capacity is automatically increased to a number)\n", cap(nums))
	fmt.Println("=========================")

	// slice with make
	numbers := make([]int, 3, 3)
	fmt.Println("[created slice with make with defined length 3 and capacity 3]")
	// length and capacity before adding
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d\n", cap(numbers))
	fmt.Println("=========================")

	// adding (assigning)
	fmt.Println("[assgined values to index 0, 1 and 2]")
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	// length and capacity after adding
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d\n", cap(numbers))
	fmt.Println("=========================")

	// adding one more element
	fmt.Println("[assgined values to index 3]")
	numbers = append(numbers, 4)
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d (*notice that capacity is automatically increased to a number)\n", cap(numbers))
	fmt.Println("=========================")
}
