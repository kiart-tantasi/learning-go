package main

import "fmt"

func main() {
	// slice
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("[slice]")
	// length and capacity before adding
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d\n", cap(nums))
	// adding
	nums = append(nums, 7)
	// length and capacity after adding
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d\n", cap(nums))

	// make slice
	numbers := make([]int, 3, 3)
	fmt.Println("[make slice]")
	// length and capacity before adding
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d\n", cap(numbers))
	// adding (assigning)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	// length and capacity after adding
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d\n", cap(numbers))
	// adding one more element
	numbers = append(numbers, 4)
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d\n", cap(numbers))
}
