package main

import "fmt"

// issue: https://stackoverflow.com/a/67422367/21331113
// in short: range over terminates only when channel is closed

func main() {
	// noDeadlock()
	// hasDeadlock()
}

func noDeadlock() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	// alternatively, you can use this classic for-loop
	for i := 0; i < 1; i++ {
		fmt.Println(<-ch)
	}
}

func hasDeadlock() {
	ch := make(chan int)
	go func() {
		// defer close(ch) // <-- one of the ways to fix this issue
		ch <- 1
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
