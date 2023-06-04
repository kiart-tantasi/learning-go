package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Queues []int
}

func (q *Queue) Enqueue(newItem int) {
	q.Queues = append(q.Queues, newItem)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.Queues) == 0 {
		return 0, errors.New("no queues left")
	}
	out := q.Queues[0]
	q.Queues = q.Queues[1:]
	return out, nil
}

func main() {
	queue := Queue{
		Queues: []int{},
	}

	// enqueue
	fmt.Println("before enqueue:", queue)
	queue.Enqueue(1)
	fmt.Println("after enqueue:", queue)
	queue.Enqueue(2)
	fmt.Println("after enqueue:", queue)
	queue.Enqueue(3)
	fmt.Println("after enqueue:", queue)

	// dequeue
	queue.Dequeue()
	fmt.Println("after dequeue:", queue)
	queue.Dequeue()
	fmt.Println("after dequeue:", queue)
	queue.Dequeue()
	fmt.Println("after dequeue:", queue)

	// no queues left
	if _, err := queue.Dequeue(); err != nil {
		fmt.Println(err)
	}
}
