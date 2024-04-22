package main


import "fmt"
type Queue struct{
	item []interface{}
}

// enqueue adds an element to the end of the queue.
func (q *Queue) enqueue(element interface{}){

	   q.item = append(q.item, element )

}

// dequeue removes and returns the element at the front of the queue.
func (q *Queue) dqueue() interface{}{

	if (len(q.item) == 0){
		return nil
	}
	q.item =q.item[1:] 
	return q.item
	
}

// func (q *Queue) front() interface{} {

// }

// func (q *Queue) rear() interface{} {

// }

func (q *Queue) isEmpty() bool{

	return true

}

func main(){

	q  := Queue{}

	// Push some elements into queue

	q.enqueue(30) 
	q.enqueue(40) 
	q.enqueue(50)

	fmt.Println(q.item)

	// Pop some elements from the queue

	q.dqueue()
	q.dqueue()
	q.dqueue()
	q.dqueue()
	//q.dqueue()
	fmt.Println(q.item)

}