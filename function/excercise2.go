package function

//https://play.golang.org/p/6X1aa8L-MZQ

import (
	"fmt"
)

func Excercise2() {

	foo := foo()
	s, i := bar()

	fmt.Println("foo output", foo)

	fmt.Println("bar ouput", s, i)
}

func foo() int {

	return 10

}

func bar() (string, int) {

	return "hello", 10

}
