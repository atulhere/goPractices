package programs

import "fmt"

func SwapTwoNumbers(a int16, b int16) {

	var c int16

	fmt.Println("The value of a and b before swap are", a, b)
	c = a
	a = b
	b = c

	fmt.Println("The value of a and b after swap are", a, b)

}
