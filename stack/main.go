package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []int
}

func (s *Stack) Push(newItem int) {
	s.Items = append(s.Items, newItem)
}

func (s *Stack) Pop() (int, error) {
	if (len(s.Items)) == 0 {
		return 0, errors.New("there are no items left in sack")
	}
	poped := s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return poped, nil
}

func main() {
	stack := Stack{
		Items: []int{},
	}

	// push
	fmt.Println("before push:", stack)
	stack.Push(1)
	fmt.Println("after push:", stack)
	stack.Push(2)
	fmt.Println("after push:", stack)
	stack.Push(3)
	fmt.Println("after push:", stack)

	// pop
	stack.Pop()
	fmt.Println("after pop:", stack)
	stack.Pop()
	fmt.Println("after pop:", stack)
	stack.Pop()
	fmt.Println("after pop:", stack)

	// no items left
	if _, err := stack.Pop(); err != nil {
		fmt.Println(err)
	}
}
