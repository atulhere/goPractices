package function

import "fmt"

//function example

//func (r reciver) identifier(parammeters) (return(s))
func SumOfNummbers(x int, y int) {

	fmt.Println("Sum of two numbers are", x+y)
}

//Varidatic parameters
func VariadicParameters(x ...int) {

	sum := 0

	for _, v := range x {

		sum = sum + v

	}

	fmt.Println("Sum of all given Numbers are", sum)
}

//Unfurling Slice

func UnfurlingSlice(x ...int) {

	sum := 0

	for _, v := range x {

		sum = sum + v

	}
	fmt.Println("the sum is", sum)
}

func ExampleAnonymousFunction() {

	func(s string) {

		fmt.Println("Hello ", s)
	}("Maa")
}
