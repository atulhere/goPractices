// Implement Linkedlist
package main

import "fmt"

//Structure for the node 
type Node struct {
	data int
	next *Node
}

// Structure for LinkedList
type LinkedList struct {
	head *Node
}

func main() {

	ll := LinkedList{}
	//insert some elements at the end of the list
	ll.insertAtTheEnd(20)
	ll.insertAtTheEnd(30)
	ll.insertAtTheEnd(40)
	ll.insertAtTheEnd(50)

	//insert an element at the begining
	ll.insertAtBegining(10)

	//insert an element after an element
	ll.insertAfterValue(30,60)

	ll.display()

}

// insertAtBegining will insert a node at the begining of the lisnked list
func  (ll *LinkedList) insertAtBegining(num int){
	     node := &Node{data:num,next:nil}
		// check if the list is empty, i.e ll.head is pointing to nil
		 if ll.head ==nil{
			 ll.head =node
			 return 
		 }
		 
		 node.next = ll.head //first point node next to head, this i to use head value
		 ll.head = node     // head will point to given node  

}

//insertAfterValue will insert a node after a given node value

func (ll *LinkedList) insertAfterValue(afterValue int, num int){

	node := &Node{data:num, next:nil}
	
	ptr := ll.head
	
	for ptr != nil{
		if(ptr.data == afterValue){	
			node.next =ptr.next
			ptr.next = node
			return 
		}
		ptr = ptr.next
	}
	
   fmt.Println("After values does not exit in the Linkedlist!")	   	
}

// insertAtTheEnd will insert the node at the end of the List

func (ll *LinkedList) insertAtTheEnd(num int) {

	tempNode := &Node{data: num}
	if ll.head == nil {
		ll.head = tempNode
		fmt.Println("Node inserted successfully ", tempNode)
		return

	}
	ptr := ll.head
	for ptr.next != nil {
		ptr = ptr.next

	}
	ptr.next = tempNode

}

// display will display all nodes of the linkedlist
func (ll *LinkedList) display() {
	ptr := ll.head

	for ptr != nil {

		fmt.Println(ptr.data)
		ptr = ptr.next

	}
}
