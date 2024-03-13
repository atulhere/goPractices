package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert a node

func (n *Node) Insert(k int) {

	// if the element is greater than root node, move to right
	// If the element is lower than the root ndoe, move to left
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}

	} else {
		if k < n.Key {
			if n.Left == nil {
				n.Left = &Node{Key: k}
			} else {
				n.Left.Insert(k)
			}

		}
	}
}
func main() {

	n := Node{Key: 10}

	n.Insert(30)
	n.Insert(5)
	fmt.Println(n)

}
