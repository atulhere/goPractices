package structure

import "fmt"

type Person struct {
	name string
	age  int
}

func ExampleStruct() {

	//https://play.golang.org/p/yp2Ht2W_DTW

	person := Person{

		name: "Atul",
		age:  50,
	}

	fmt.Println("The name and age of the Persion are", person.name, person.age)
}
