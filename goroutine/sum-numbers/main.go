// to try to use for loop to sum numbers in slice with / without goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = []int{10, 10, 10, 10, 10}
	var slice3 = []int{1000, 1000, 1000, 1000, 1000}

	// sum without goroutine
	start := time.Now()
	final1 := sum(slice1)
	final2 := sum(slice2)
	final3 := sum(slice3)
	fmt.Println("finished in", time.Since(start))
	fmt.Println("results:", final1, final2, final3)
	fmt.Println("")

	// sum with goroutine
	start = time.Now()
	result1 := make(chan int)
	result2 := make(chan int)
	result3 := make(chan int)
	go sumGoroutine(slice1, result1)
	go sumGoroutine(slice2, result2)
	go sumGoroutine(slice3, result3)
	final1 = <-result1
	final2 = <-result2
	final3 = <-result3
	fmt.Println("finished in", time.Since(start))
	fmt.Println("results:", final1, final2, final3)
}

func sum(slice []int) int {
	time.Sleep(2 * time.Second) // simulate slow response
	var sum int
	for _, val := range slice {
		sum += val
	}
	return sum
}

func sumGoroutine(slice []int, result chan int) {
	var sumResult = sum(slice)
	result <- sumResult
}
