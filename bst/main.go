package main

import "fmt"

type Bst struct {
	headNode *Node
}

func (b *Bst) Insert(newValue int) {
	if b.headNode == nil {
		b.headNode = &Node{
			value: newValue,
		}
	} else {
		b.headNode.Insert(newValue)
	}
}

func (b *Bst) Search(value int) int {
	if b.headNode == nil {
		return -1
	}
	return b.headNode.Search(value)
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (node *Node) Insert(newValue int) {
	if newValue <= node.value {
		// left
		if node.left == nil {
			node.left = &Node{
				value: newValue,
			}
		} else {
			node.left.Insert(newValue)
		}
	} else {
		// right
		if node.right == nil {
			node.right = &Node{
				value: newValue,
			}
		} else {
			node.right.Insert(newValue)
		}
	}
}

func (node *Node) DeptFirstTraverse() {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	node.left.DeptFirstTraverse()
	node.right.DeptFirstTraverse()
}

func (node *Node) Search(value int) int {
	if node == nil {
		return -1
	}
	if node.value == value {
		return node.value
	}
	if value < node.value {
		return node.left.Search(value)
	} else {
		return node.right.Search(value)
	}
}

// TODO: make fimction to do breadth-first traverse

func main() {
	bst := &Bst{}
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)
	bst.Insert(5)
	bst.headNode.DeptFirstTraverse()

	// search
	fmt.Println("\nsearching")
	fmt.Println(bst.Search(3))
	fmt.Println(bst.Search(2))
	fmt.Println(bst.Search(4))
	// search for node that does not exist
	fmt.Println(bst.Search(123))
}

//       3
//     2   4
//    1     5
