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

// Pop an element from the slice
func (s *Stack) pop(){

	l := len(s.item) -1 
	s.item = s.item[:l]

}

// Check If Stack is empty

func (s *Stack) isEmpty() string{
	if(len(s.item)>0){
		return "Stack is non empty"
	}
	return "Stack is empty"
}

// Print the top element from the stack
func (s * Stack) top(){

	if(len(s.item)>0){
		 element := s.item[len(s.item)-1]
		 fmt.Println("Stack Top element is : ", element)
	}else{

		fmt.Println("Stack is empty")
	}	
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

	// Check if Stack is empty
	response := s.isEmpty()
	fmt.Println(response)

	// Print the top element from the stack
	s.top()
}




