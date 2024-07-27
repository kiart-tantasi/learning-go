package main

import (
	"fmt"
	"time"
)

// [from chatgpt]
// If your function only needs to receive data from the channel and you want to enforce this constraint,
// use <-chan int. This makes your code more self-documenting and helps prevent accidental misuse of the channel.
// If your function needs to send data on the channel, use chan int.

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	go func() {
		canDoBoth(ch)
	}()

	go func() {
		receiveOnly(ch)
	}()

	// hack
	time.Sleep(1 * time.Second)
}

func receiveOnly(ch <-chan int) {
	// ch <- 2 // invalid operation: cannot send to receive-only channel ch (variable of type <-chan int)compilerInvalidSend
	for val := range ch {
		fmt.Println("receive only:", val) // or "<- ch"
	}
}

func canDoBoth(ch chan int) {
	ch <- 2
	// fmt.Println("both:", <-ch)
}
