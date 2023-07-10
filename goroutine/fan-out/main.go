// partially based on course "Go: Data Structures, Algorithms and Design Patterns with Go | Udemy"
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// producers
	producer := make(chan int)
	go produce(producer)
	// consumer
	consumer1 := make(chan int)
	consumer2 := make(chan int)
	go consume(consumer1, "1")
	go consume(consumer2, "2")
	// fan-out
	fanOut(producer, consumer1, consumer2)
}

func delay() {
	time.Sleep(1 * time.Second)
}

func produce(ch chan int) {
	for {
		delay()
		random := rand.Intn(10)
		ch <- random
	}
}

func consume(ch <-chan int, id string) {
	for num := range ch {
		fmt.Printf("consumer %s consumed %d\n", id, num)
	}
}

func fanOut(producer <-chan int, consumer1, consumer2 chan<- int) {
	i := 0
	for msg := range producer {
		if i%2 == 0 {
			consumer1 <- msg
		} else {
			consumer2 <- msg
		}
		i++
	}
}
