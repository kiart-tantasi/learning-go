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

func deadlockDemo() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	go func() {
		time.Sleep(2 * time.Second)
		// free space
		<-ch // un-comment this line to resolve deadlock (channel blocking but no code will free space)
	}()

	ch <- 4

	fmt.Println("poped from channel -", (<-ch))
	fmt.Println("poped from channel -", (<-ch))
	fmt.Println("poped from channel -", (<-ch))
	fmt.Print("done")
}
