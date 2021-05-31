package structure

//https://play.golang.org/p/6ihcMzaXp-7

import (
	"fmt"
)

type detail struct {
	fname           string
	lname           string
	icecreamFlavour []string
}

func MapThroughStruct() {

	person1 := detail{fname: "Atul", lname: "Singh", icecreamFlavour: []string{"chocolate", "kulfi"}}

	person2 := detail{fname: "Priti", lname: "Singh", icecreamFlavour: []string{"Vanila", "mango"}}

	m := map[string]detail{
		person1.fname: person1,
		person2.fname: person2,
	}

	fmt.Println(m)
	for i, v := range m {

		fmt.Println(i, v)

	}
}
