package main

import (
	"fmt"
	"time"
)

// TODO: use context package to help cancel unused goroutine

func main() {
	channel := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("goroutine is finished")
		channel <- "value"
	}()

	select {
	case val := <-channel:
		fmt.Println(val)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
		time.Sleep(3 * time.Second)
		return
	}
}
