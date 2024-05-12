package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	demoSelect()
	// demoResourceLeak()
}

func demoSelect() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	// functions to send values into channels
	go goRoutineOne(chan1)
	go goRoutineTwo(chan2)
	// use select with for-loop to receive values from channel
	for {
		select {
		case val := <-chan1:
			fmt.Println(val)
		case val := <-chan2:
			fmt.Println(val)
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			os.Exit(0)
		}
	}
}

// to demo that if we just use normal timeout, goroutine will still continue and consume resource
func demoResourceLeak() {
	channel := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("goroutine is finished")
		channel <- 1
	}()

LOOP:
	for {
		select {
		case val := <-channel:
			fmt.Println(val)
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
			// wait for goroutine to finish to prove that it is still running
			time.Sleep(5 * time.Second)
			break LOOP
		}
	}
	fmt.Println("closing app")
}

func goRoutineOne(chan1 chan string) {
	time.Sleep(time.Duration(rand.Intn(3) * int(time.Second)))
	chan1 <- "from goRoutineOne"
}

func goRoutineTwo(chan2 chan string) {
	time.Sleep(time.Duration(rand.Intn(3) * int(time.Second)))
	chan2 <- "from goRoutineTwo"
}
