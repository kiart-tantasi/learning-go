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
	fmt.Println("started app (doesWait)")
	channal := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep is done")
		channal <- 123
	}()
	valueFromChannel := <-channal
	fmt.Println("value is", valueFromChannel)
	fmt.Println("app ended")
	fmt.Println("")
}

func doesNotWait() {
	fmt.Println("started app (doesNotWait)")
	channal := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep is done")
		channal <- 123
	}()
	// value in channel is not pulled out and used
	fmt.Println("app ended")
	fmt.Println("")
}
