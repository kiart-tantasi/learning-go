package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (l *DoublyLinkedList) Add(node *Node) {
	if l.Head == nil {
		l.Head = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
	}
	l.Tail = node
}

func (l *DoublyLinkedList) PrintAllNodesValue() {
	node := l.Head
	for node != nil {
		fmt.Println("current node's value:", node.Value)
		node = node.Next
	}
}

func (l *DoublyLinkedList) PrintAllNodesValueBackWard() {
	node := l.Tail
	for node != nil {
		fmt.Println("current node's value:", node.Value)
		node = node.Prev
	}
}

func (l *DoublyLinkedList) FindIndex(value int) (int, error) {
	index := 0
	node := l.Head

	for node != nil {
		if node.Value == value {
			return index, nil
		}
		index++
		node = node.Next
	}

	return 1, errors.New("node not found")
}

func (l *DoublyLinkedList) getLength() int {
	length := 0
	node := l.Head
	for node != nil {
		node = node.Next
		length++
	}

	return length
}

func main() {
	list := &DoublyLinkedList{}
	for i := 1; i <= 3; i++ {
		list.Add(&Node{
			Value: i,
		})
	}
	list.PrintAllNodesValue()
	fmt.Println("")

	list.PrintAllNodesValueBackWard()
	fmt.Println("")

	valueToFind := 1
	if index, err := list.FindIndex(valueToFind); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("node index:", index)
	}
	fmt.Println("")

	fmt.Print("length of list:", list.getLength())
	fmt.Println("")
}
