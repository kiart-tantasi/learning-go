package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// produce 1
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}(ch1)

	// produce 2
	go func(ch chan int) {
		for i := 11; i < 20; i++ {
			ch <- i
		}
	}(ch2)

	// consue1
	consumer := make(chan int)
	go func(c <-chan int) {
		for val := range c {
			fmt.Println("consumer got value", val)
		}
	}(consumer)

	// combine producers and consumer
	for {
		select {
		case val := <-ch1:
			consumer <- val
		case val := <-ch2:
			consumer <- val
		default:
			fmt.Println("sleeping...")
			time.Sleep(1 * time.Second)
		}
	}
}
