package main

import (
	"fmt"
)

func main() {
	q := make(chan int, 4)
	// enqueue
	fmt.Println("starting enqueueing")
	q <- 1
	q <- 2
	q <- 3
	q <- 4
	// q <- 5 // will cause error "fatal error: all goroutines are asleep - deadlock!"
	fmt.Println(("finished enqueueing"))

	// dequeue
	fmt.Println("starting dequeueing")
	fmt.Println(<-q)
	fmt.Println(<-q)
	fmt.Println(<-q)
	fmt.Println(<-q)
	// fmt.Println(<-q) // will cause error "fatal error: all goroutines are asleep - deadlock!"
	fmt.Println(("finished dequeueing"))
}
