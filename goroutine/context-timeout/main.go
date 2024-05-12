package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// demoWithTimeout()
	demoWithCancel()

	// context.WithDeadline works pretty close to .WithTimeout except we can instead define expiration time
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

func demoWithCancel() {
	isJobDone := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // prevent resource leak

	// call CancelFunc after 2 seconds
	go func(cancel context.CancelFunc) {
		time.Sleep(2 * time.Second)
		cancel()
	}(cancel)

	// job will be done in 3 seconds which is after calling context.CancelFunc
	// you can change to 1 second to make job done before
	go func(channel chan bool) {
		time.Sleep(3 * time.Second)
		isJobDone <- true
	}(isJobDone)

	// let's check result
	channel := make(chan string)
	go func(ctx context.Context, channel chan string) {
		for {
			select {
			case <-ctx.Done():
				channel <- ctx.Err().Error()
				return
			case <-isJobDone:
				channel <- "job is done before CancelFunc, yeahhh!"
				return
			}
		}
	}(ctx, channel)
	fmt.Println("result:", <-channel)
}
