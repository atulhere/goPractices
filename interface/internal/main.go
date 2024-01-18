package main

import "fmt"

type Employee interface {
	salary() int
	retire()
}

type Worker struct {
	sal int
}

func (worker Worker) salary() int {

	return worker.sal
}

func (worker Worker) retire() {

	fmt.Println("I am retiring")
}

func main() {

	var emp Employee // Here Type and Value of interface is NIL

	w := Worker{30000}

	emp = w

	fmt.Printf("The type of inteface %T and value%v\n", emp, emp)
	sal := emp.salary()
	fmt.Println("The salary of employee is", sal)

}
