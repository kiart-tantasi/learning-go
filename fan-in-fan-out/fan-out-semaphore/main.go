package main

import (
	"fmt"
	"time"
)

// summary: 2 consumers have their own semaphore of 2 goroutines

func main() {
	// produce 1
	producer := make(chan int)
	go func(ch chan int) {
		defer close(ch)
		for i := 0; i < 10; i++ {
			producer <- i
		}
	}(producer)

	consumer1Semaphore := make(chan int, 2)
	consumer2Semaphore := make(chan int, 2)
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
					fmt.Println("consumer 1, before reserving semaphore for value:", val)

					// reserve semaphore
					consumer1Semaphore <- 1

					fmt.Println("consumer 1, started, value:", val)

					// pretend this consumer takes long tasks
					time.Sleep(5 * time.Second)

					// release semaphore
					<-consumer1Semaphore
				}()
			} else {
				// consumer 2
				go func() {
					fmt.Println("consumer 2, before reserving semaphore for value:", val)

					// reserve semaphore
					consumer2Semaphore <- 1

					// log before starting
					fmt.Println("consumer 2, started, value:", val)

					// pretend this consumer takes short tasks
					time.Sleep(2 * time.Second)

					// release semaphore
					<-consumer2Semaphore
				}()
			}
		default:
			// no values found at the moment
		}
	}
}

func isConsumer1(val int) bool {
	// get all odd numbers
	return val%2 != 0
}
