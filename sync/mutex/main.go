package main

import (
	"fmt"
	"sync"
)

func main() {
	demoWithMutex()
	demoWithoutMutex()
}

func demoWithoutMutex() {
	numGoroutines := 10
	var wg sync.WaitGroup
	var counter int
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementCounter(&wg, &counter)
	}
	wg.Wait()
	fmt.Println("without mutex:", counter)
}

func demoWithMutex() {
	numGoroutines := 10
	var wg sync.WaitGroup
	var counter int
	var mutex sync.Mutex
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementCounterMutex(&wg, &counter, &mutex)
	}
	wg.Wait()
	fmt.Println("with mutex:", counter)
}

func incrementCounter(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	for i := 0; i < 100_000; i++ {
		*counter++
	}
}

// NOTE: value-receiver of sync.Mutex does not work, only pointer-receiver does ***
func incrementCounterMutex(wg *sync.WaitGroup, counter *int, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	for i := 0; i < 100_000; i++ {
		*counter++
	}
	mutex.Unlock()
}
