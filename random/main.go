package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomInt := rand.Intn(3000)
	fmt.Println("value:", randomInt)

	start := time.Now()
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
	fmt.Printf("slept for %d milliseconds\n", time.Since(start))
}
