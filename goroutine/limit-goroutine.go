package main

import (
	"fmt"
	"time"
)

func main() {
	resultAmount := 9

	limiter := make(chan int, 3)
	results := make(chan string, resultAmount)
	for i := 1; i <= resultAmount; i++ {
		limiter <- 0
		go doTask(i, results, limiter)
	}
	for i := 1; i <= resultAmount; i++ {
		fmt.Println("result:", <-results)
	}
}

func doTask(taskNum int, results chan string, limiter chan int) {
	fmt.Println("started task", taskNum)
	time.Sleep(3 * time.Second)
	results <- fmt.Sprintf("hello world %d", taskNum)
	fmt.Println("finished task", taskNum)
	<-limiter
}
