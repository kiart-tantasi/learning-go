package main

import (
	"context"
	"fmt"
)

type ContextKey string

const (
	ContextKey1 ContextKey = "ContextKey1"
	ContextKey2 ContextKey = "ContextKey2"
)

func main() {
	// ctx2 would have value1 from parent context (ctx1) and also have newly added value2
	ctx1 := context.WithValue(context.Background(), ContextKey1, "value1")
	ctx2 := context.WithValue(ctx1, ContextKey2, "value2")

	// ctx1
	fmt.Println("ctx1's ContextKey1:", ctx1.Value(ContextKey1))
	fmt.Println("ctx1's ContextKey2:", ctx1.Value(ContextKey2))

	// ctx2
	fmt.Println() // for readability sake
	fmt.Println("ctx2's ContextKey1:", ctx2.Value(ContextKey1))
	fmt.Println("ctx2's ContextKey2:", ctx2.Value(ContextKey2))
}
