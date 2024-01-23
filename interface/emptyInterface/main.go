package main

import "fmt"

type EmptyInterface interface{}

func main() {

	var emp EmptyInterface

	emp = "Hello World!"

	fmt.Println("Empty interface having String Value", emp)

	emp = 1000

	fmt.Println("Empty interface having Integer Value", emp)

}
