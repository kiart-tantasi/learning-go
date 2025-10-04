package main

import "fmt"

// issue: https://stackoverflow.com/a/67422367/21331113
// in short: range-over terminates only when channel is closed

func main() {
	// hasDeadlock()
	noDeadlock1()
	// noDeadlock2()
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

func noDeadlock1() {
	ch := make(chan int)
	go func() {
		ch <- 1
		defer close(ch) // fix (close channel so range-over does not get deadlock)
	}()
	for val := range ch {
		fmt.Println(val)
	}

	// check channel status after it is closed
	if _, ok := <-ch; ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed")
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
