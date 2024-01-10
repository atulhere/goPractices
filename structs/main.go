package main

import "fmt"

// Declare a struct type
type Person struct {
	name string
	age  int
	city string
}

type Employee struct {
	p       Person //Example of Embeded Struct
	salary  int
	company string
}

func main() {

	//declare a variable of Person type(Which is an instance of Person)
	var person Person
	fmt.Println("The default valaue of the person is ", person)

	//lets assign some values to the instance person
	person = Person{"Atul", 38, "Jaunpur"}
	fmt.Println("The valaue of the person is ", person)

	/*
		let's create another instance of Person
		Struct and assign some values using Short
		declaration operator
	*/
	p1 := Person{"Rahul", 39, "Azamgarh"}
	//Note:There is no need to put field keys if we are adding values in order

	fmt.Println("The value of the stuct is ", p1)

	//Example for Embeded struct
	embS := Employee{

		p: Person{
			name: "Atul",
			age:  38,
			city: "Jaunpur",
		},
		salary:  10000,
		company: "XYZ",
	}

	fmt.Println("Example of Embeded Struct is ", embS)

}
