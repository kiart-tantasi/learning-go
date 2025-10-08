package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	stopCh := make(chan int)

	go func() {
		time.Sleep(200 * time.Millisecond)
		close(stopCh)
	}()

	for {
		select {
		case val := <-stopCh:
			fmt.Printf("got value from stop channel [%d]\n", val)
			os.Exit(0)
		default:
			fmt.Println("no value from stop channel")
		}
	}
}
