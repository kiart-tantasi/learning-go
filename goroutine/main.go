package main

import (
	"fmt"
	"time"
)

func main() {
	// doesWait()
	// doesNotWait()
	// deadlockDemo()
}

func deadlockDemo() {
	size := 3
	limited := make(chan int, size)
	limited <- 1
	limited <- 2
	limited <- 3
	// limited <- 4 // causes deadlock error (exceeding size)

	toAdd := 4
	for {
		// pull value from channel
		fmt.Println(<-limited)

		// keep adding value into channel to prevent deadlock error (no value to be pulled)
		limited <- toAdd // comment this to get deadlock error
		toAdd++
		time.Sleep(time.Second)
	}
}

func doesWait() {
	fmt.Println("[doesWait]")
	channal := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("sleep is done")
		channal <- 123
	}()
	// pull value from channel to block the next line
	<-channal
	fmt.Println("app ended\n")
}

func doesNotWait() {
	fmt.Println("[doesNotWait]")
	channal := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("sleep is done")
		channal <- 123
	}()
	// value in channel is not pulled out so it does not block the next line
	fmt.Println("app ended\n")
}
