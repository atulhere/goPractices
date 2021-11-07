package method

import (
	"fmt"
)

type employee struct {
	name string
	age  int
}

func ExmapleMethod() {

	emp1 := employee{

		name: "atul",
		age:  36,
	}
	emp2 := employee{

		name: "priti",
		age:  33,
	}

	fmt.Println("Example of Methods")
	emp1.employeeInfo()

	emp2.employeeInfo()
}

func (em employee) employeeInfo() {

	fmt.Println(em.name)

}
