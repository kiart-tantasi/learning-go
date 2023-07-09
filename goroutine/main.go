package main

import (
	"fmt"
	"time"
)

func main() {
	doesWait()
	doesNotWait()
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
