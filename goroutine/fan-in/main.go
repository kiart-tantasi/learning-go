// partially based on course "Go: Data Structures, Algorithms and Design Patterns with Go | Udemy"
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// producers
	producer1 := make(chan int)
	producer2 := make(chan int)
	go produce(producer1, "1")
	go produce(producer2, "2")
	// consumer
	consumer := make(chan int)
	go consume(consumer)
	// fan-in function
	fanOut(producer1, producer2, consumer)
}

func delay() {
	time.Sleep(1 * time.Second)
}

func produce(ch chan int, id string) {
	for {
		delay()
		random := rand.Intn(10)
		fmt.Printf("producer %s produced %d\n", id, random)
		ch <- random
	}
}

func consume(ch <-chan int) {
	for num := range ch {
		fmt.Printf("consumer consumed %d\n", num)
	}
}

func fanOut(producer1, producer2 <-chan int, consumer chan<- int) {
	for {
		select {
		case receiver := <-producer1:
			consumer <- receiver
		case receiver := <-producer2:
			consumer <- receiver
		default:
			fmt.Println("no messages found, will retry in 1 second")
			time.Sleep(time.Second)
		}
	}
}

// === testConsumeFunction consume function === //
// func testConsumeFunction() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 1
// 		ch <- 2
// 		ch <- 3
// 	}()
// 	go consumeFunction(ch)
// 	time.Sleep(1 * time.Millisecond) // hack
// }
// func consumeFunction(ch <-chan int) {
// 	for msg := range ch {
// 		fmt.Println("Consumed:", msg)
// 	}
// }
