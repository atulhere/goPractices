package main

import "fmt"

// A simple program to understand the use of Interface
type Employee interface {
	salary() int
}

type Emp struct {
	sal   int
	grade string
}

func (emp Emp) salary() int {
	pf := 1800
	return emp.sal + pf

}

type Animal struct {
	message string
}

type Cat struct {
	message string
}

func (animal Animal) speak() {

	fmt.Println("Hello from ", animal.message)

}
func (animal Cat) speak() {

	fmt.Println("Hello from ", animal.message)

}

func main() {

	var result int
	//var obj Employee = Emp{15000, "A"}

	var obj Employee
	obj = Emp{2000, "B"}

	result = obj.salary()
	fmt.Println("Salary is ", result)

	// Assigning some values to Animal struct
	dog := Animal{"Dog"}
	dog.speak()

	// Assigning some values to Cat struct

	cat := Cat{"Cat..."}
	cat.speak()

}
