package datatype

import "fmt"

func PracticeType() {

	var a = 10

	type atul int

	var b atul

	a = int(b)

	fmt.Printf("The value of b is %v and type is %T\n", b, b)
	fmt.Printf("The value of a is %v and type is %T\n", a, a)

}
