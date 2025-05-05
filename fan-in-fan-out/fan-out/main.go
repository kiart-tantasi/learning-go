package main

import (
	"fmt"
)

// summary: single producer with multiple consumers

func main() {
	// produce 1
	producer := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			producer <- i
		}
		close(ch)
	}(producer)

	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// use for-loop to distribute values from a single channel to multiple goroutines
	run := true
	for run {
		select {
		case val, ok := <-producer:
			// skip if retrieving invalid value from channel
			if !ok {
				continue
			}

			if isConsumer1(val) {
				// consumer 1
				go func() {
					fmt.Println("consumer 1's value:", val)
					// ch1 <- val
				}()
			} else {
				// consumer 2
				go func() {
					fmt.Println("consumer 2's value:", val)
					// ch2 <- val
				}()
			}

		default:
			// no values found at the moment
		}
	}

	// we can also do fan-in from ch1 and ch2 here (code for ch1 and ch2 is commented aboce)
}

func isConsumer1(val int) bool {
	// get all odd numbers
	return val%2 != 0
}
