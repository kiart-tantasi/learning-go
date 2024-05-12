package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	demoWithTimeout()
}

func demoWithTimeout() {
	channel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // prevent resource leak

	go func(ctx context.Context, channel chan string) {
		for {
			select {
			// timeout
			case <-ctx.Done():
				channel <- ctx.Err().Error()
				return
			// pretend that job is done after 5 seconds which already exceeds timeout
			// you can set it to 1 second so the job is done before timeout
			case <-time.After(5 * time.Second):
				channel <- "job is done before timeout, yeahhh!"
				return
			}
		}
	}(ctx, channel)

	// let's check result
	fmt.Println("result:", <-channel)
}
