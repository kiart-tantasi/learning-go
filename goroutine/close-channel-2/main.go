package main

import (
	"fmt"
	"time"
)

// when trying to receive value from closed channel, you will get 0 (default value)
// the way to avoid it is to use syntax "val,ok := <- ch" or to leave channel opened

func main() {
	getValueFromClosedChannel()
}

func getValueFromClosedChannel() {
	ch := make(chan int)
	go func() {
		defer close(ch) // comment this line and you will see the difference (we will not get values from being-opened channel)
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for {
		select {
		// get a value
		case val := <-ch:
			fmt.Println("val:", val)
		// no value to get
		default:
			fmt.Println("sleeping...")
			time.Sleep(1 * time.Second)
		}
	}
}
