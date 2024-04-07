package main

import "fmt"


// Stack represent a stack which hold a slice of item
type Stack struct{

	item []int

}

// Push an element into the Slice

func (s *Stack) push(number int){

	s.item = append(s.item, number)

}

// pop an element from the slice
func (s *Stack) pop(){

	l := len(s.item) -1 
	s.item = s.item[:l]

}


func main(){

	// Create a object of Stack type
	s := Stack{} 

	// Push some some integer numbers into Stack
	s.push(20)
	s.push(30)
	s.push(50)
	s.push(70)
	s.push(90)

	// Print the elements
	fmt.Println(s.item)

	// Pop the last element from the stack
	s.pop()

	// Print the elements
	fmt.Println(s.item)

}




