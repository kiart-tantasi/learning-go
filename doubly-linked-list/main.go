package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (l *DoublyLinkedList) Add(node *Node) {
	if l.Head == nil {
		l.Head = node
	} else {
		node.prev = l.Tail
		l.Tail.next = node
	}
	l.Tail = node
}

func (l *DoublyLinkedList) PrintAllNodesValue() {
	node := l.Head
	for node != nil {
		fmt.Println("current node's value:", node.value)
		node = node.next
	}
}

func (l *DoublyLinkedList) PrintAllNodesValueBackWard() {
	node := l.Tail
	for node != nil {
		fmt.Println("current node's value:", node.value)
		node = node.prev
	}
}

func (l *DoublyLinkedList) FindIndex(value int) (int, error) {
	index := 0
	node := l.Head

	for node != nil {
		if node.value == value {
			return index, nil
		}
		index++
		node = node.next
	}

	return 1, errors.New("node not found")
}

func (l *DoublyLinkedList) getLength() int {
	length := 0
	node := l.Head
	for node != nil {
		node = node.next
		length++
	}

	return length
}

func main() {
	list := &DoublyLinkedList{}
	for i := 1; i <= 3; i++ {
		list.Add(&Node{
			value: i,
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
