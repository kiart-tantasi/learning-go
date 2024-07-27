package main

import "fmt"

func main() {
	// produce 1
	producer := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			producer <- i
		}
		close(ch)
	}(producer)

	switcher := false
	ch1 := make(chan int)
	ch2 := make(chan int)
	for val := range producer {
		if switcher {
			// consumer 1
			go func() {
				fmt.Println("consumer 1's value:", val)
				ch1 <- val
			}()
		} else {
			// consumer 2
			go func() {
				fmt.Println("consumer 2's value:", val)
				ch2 <- val
			}()
		}
		// change switch
		switcher = !switcher
	}

	// we can also do fan-in from ch1 and ch2 here
}
