package main

import (
	"fmt"
	"time"
)

func main() {
	doesWait()
	doesNotWait()
	getValueFromGoroutine()
}

func doesWait() {
	fmt.Println("started app (doesWait)")
	channal := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep is done")
		channal <- 123
	}()
	<-channal
	fmt.Println("app ended\n")
}

func doesNotWait() {
	fmt.Println("started app (doesNotWait)")
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep is done")
	}()
	fmt.Println("app ended\n")
}

func getValueFromGoroutine() {
	fmt.Println("started app (getValueFromGoroutine)")
	strChan := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		strChan <- "Hello World"
	}()
	// get value from channal variable
	strFromChan := <-strChan
	fmt.Println(strFromChan)
	fmt.Println("app ended\n")
}
