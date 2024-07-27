package main

import "fmt"

// issue: https://stackoverflow.com/a/67422367/21331113
// in short: range over terminates only when channel is closed

func main() {
	noDeadlock1()
	noDeadlock2()

	// hasDeadlock()
}

func noDeadlock1() {
	ch := make(chan int)
	go func() {
		ch <- 1
		// close channel to prevent deadlock
		close(ch)
	}()
	for val := range ch {
		fmt.Println(val)
	}
}

func noDeadlock2() {
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
		ch <- 1
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
