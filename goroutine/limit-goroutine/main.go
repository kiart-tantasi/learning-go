package main

import (
	"fmt"
	"time"
)

func main() {
	resultAmount := 9
	limiter := make(chan int, 3) // ony allow 3 goroutines to run in parallel
	results := make(chan string, resultAmount)

	// === these 2 for-loop will run in parellel at the end because the first one is executing goroutines ===
	for i := 1; i <= resultAmount; i++ {
		// reserve the space in limiter channel to make other goroutine wait
		limiter <- 0
		go doTask(i, results, limiter)
	}
	for i := 1; i <= resultAmount; i++ {
		// pull value from result channel
		fmt.Println("result:", <-results)
	}
}

func doTask(taskNum int, results chan string, limiter chan int) {
	time.Sleep(3 * time.Second)
	results <- fmt.Sprintf("result for task no. %d", taskNum)
	fmt.Println("finished task no.", taskNum)
	// pull value from channel (free the space for next task)
	<-limiter
}
