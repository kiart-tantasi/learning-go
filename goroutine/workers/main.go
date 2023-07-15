package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	go produce(in)

	// multiple workers
	out := make(chan int)
	go worker(in, out)
	// go worker(in, out)
	// go worker(in, out)
	// go worker(in, out)

	// received message from finished processes
	for number := range out {
		fmt.Println("done:", number)
	}

}

func fakeProcess() {
	time.Sleep(3 * time.Second)
}

func produce(ch chan<- int) {
	i := 0
	for {
		ch <- i
		i++
	}
}

func worker(in <-chan int, out chan<- int) {
	for {
		fakeProcess()
		out <- <-in
	}
}
