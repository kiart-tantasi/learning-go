package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	limitDemo()
	// limitDeadlockDemo()
}

// this demo is complex but in short, goroutine has magic:
// goroutine has blocking behavior when channel is full

func limitDemo() {
	limit := 2
	limiter := make(chan int, limit)
	resultAmount := 4
	results := make(chan int, resultAmount)

	for num := 1; num <= resultAmount; num++ {
		// reserve space in limit channel
		// because we have code to free space in function below so we can reserve space in limiter channel with more then limit
		limiter <- 0
		fmt.Println("reserved limit for number", num)
		go func(currentNum int) {
			time.Sleep(5 * time.Second)
			results <- currentNum
			// free space
			<-limiter
			fmt.Println("freed space for number", currentNum)
		}(num)
	}

	getResultsCorrectDemo(resultAmount, results)
	// getResultsWrongDemo(results)
}

func limitDeadlockDemo() {
	limit := 2
	limiter := make(chan int, limit)
	resultAmount := 4
	results := make(chan string, resultAmount)

	for num := 1; num <= resultAmount; num++ {
		limiter <- 0
		currentNum := num
		go func() {
			time.Sleep(5 * time.Second)
			str := strconv.Itoa(currentNum)
			results <- str
			// we never promise to free space in limiter channel
			// so deadlock error wiil be throwned
			// <-limiter // un-comment to promise to free space to prevent deadlock error
		}()
	}
}

func getResultsCorrectDemo(resultAmount int, results chan int) {
	resultsStr := ""
	for i := 1; i <= resultAmount; i++ {
		resultsStr += strconv.Itoa(<-results) + " "
	}
	fmt.Println("results:", resultsStr)
}

func getResultsWrongDemo(results chan int) {
	resultsStr := ""
	for i := 0; i < len(results); i++ {
		resultsStr += strconv.Itoa(<-results) + "-"
	}
	fmt.Println("results:", resultsStr)
}
