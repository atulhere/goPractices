package main

import (
	"fmt"
)

// Node represent a node in Linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represent a singly linked list
type LinkedList struct {
	Head *Node
}

// Insert insert a new node with the given data into the end of the Linked List
func (ll *LinkedList) Insert(data int) {

	newNode := &Node{Data: data}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	lastNode := ll.Head

	for lastNode.Next != nil {

		lastNode = lastNode.Next

	}
	lastNode.Next = newNode
}

func (ll *LinkedList) Display() {

	current := ll.Head

	for current != nil {

		fmt.Println(current.Data)
		current = current.Next
	}

}
func main() {

	//Create a linked list
	linkedList := LinkedList{}

	// Insert a Element( a node) in Link List

	linkedList.Insert(20)
	linkedList.Insert(30)
	linkedList.Insert(40)
	linkedList.Insert(50)
	linkedList.Insert(60)

	// Display the elements from the Link List

	linkedList.Display()

}
