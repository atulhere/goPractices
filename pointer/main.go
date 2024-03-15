package main

import "fmt"

func main() {
	a := 10
	fmt.Println("Value of a is ", a)
	//using address of a
	copy(&a)
	fmt.Println("Value of a outside the function call is ", a)

}

func copy(a *int) {

	*a = *a + 30
	fmt.Println("Updated Value of A inside another function is ", *a)
}
