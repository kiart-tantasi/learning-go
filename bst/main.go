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

func (bst *Bst) DepthFirstTraverse() {
	if bst.headNode == nil {
		return
	}
	bst.headNode.DeptFirstTraverse()
}

func (bst *Bst) BreadthFirstTraverse() {
	if bst.headNode == nil {
		return
	}
	nodes := []*Node{
		bst.headNode,
	}
	for len(nodes) != 0 {
		node := nodes[0]
		node.BreadthFirstTraverse(&nodes)
	}
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (node *Node) Insert(newValue int) {
	if newValue <= node.value {
		if node.left == nil {
			node.left = &Node{
				value: newValue,
			}
		} else {
			node.left.Insert(newValue)
		}
	} else {
		if node.right == nil {
			node.right = &Node{
				value: newValue,
			}
		} else {
			node.right.Insert(newValue)
		}
	}
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

func (node *Node) DeptFirstTraverse() {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	node.left.DeptFirstTraverse()
	node.right.DeptFirstTraverse()
}

func (node *Node) BreadthFirstTraverse(items *[]*Node) {
	if node.left != nil {
		*items = append(*items, node.left)
	}
	if node.right != nil {
		*items = append(*items, node.right)
	}
	fmt.Println(node.value)
	*items = (*items)[1:]
}

func main() {
	//   [tree]
	//         5
	//     3       7
	//   2   4   6   8
	// 1               9
	bst := &Bst{}
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(1)
	bst.Insert(9)

	// depth-first traverse
	fmt.Println("depth-first traverse:")
	bst.DepthFirstTraverse()

	// breadth-first traverse
	fmt.Println("")
	fmt.Println("breadth-first traverse:")
	bst.BreadthFirstTraverse()

	// search
	fmt.Println("")
	fmt.Println("searching")
	fmt.Println(bst.Search(3))
	fmt.Println(bst.Search(2))
	fmt.Println(bst.Search(4))
	// search for node that does not exist
	fmt.Println(bst.Search(123))
}
