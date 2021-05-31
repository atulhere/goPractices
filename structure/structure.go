package structure

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

type Salary struct {
	Person
	salary int
	tax    int
}

func ExampleStruct() {

	//https://play.golang.org/p/yp2Ht2W_DTW

	person := Person{

		name: "Atul",
		age:  50,
	}

	fmt.Println("The name and age of the Persion are", person.name, person.age)
}

func EmbededStructs() {

	salary := Salary{

		Person: Person{

			name: "Atul",

			age: 35,

			address: "Ahammadpur",
		},
		salary: 110000,
		tax:    100,
	}

	fmt.Println("Exmaple of embeded Struct is", salary)

}

func AnonymousStruct() {

	company := struct {
		name      string
		emplpoyee int
	}{
		name:      "Ayoconnect",
		emplpoyee: 150,
	}

	fmt.Println("Exmaple of Anonymous Struct", company)

}
