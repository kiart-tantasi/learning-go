package main

import (
	"fmt"
	"time"
)

// when trying to receive value from closed channel, you will get 0 (default value)
// the way to avoid it is to use syntax "val,ok := <- ch" or to leave channel opened

func main() {
	valueFromClosedChannel()
}

func valueFromClosedChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch) // comment this line and you will see the difference
	}()

	for {
		select {
		case val := <-ch:
			fmt.Println("val:", val)
		default:
			fmt.Println("sleeping...")
			time.Sleep(1 * time.Second)
		}
	}
}
