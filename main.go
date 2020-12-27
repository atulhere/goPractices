package main

import (
	"fmt"

	"goPractices/constant"
	"goPractices/hello"
	"goPractices/structure"
)

func main() {

	//Hello Word from Package
	hello.HellWorld()

	//constant Practice
	constant.PlayWithConstant()

	//Example Control Flow
	//controlflow.PrintHelloWordNtimes()

	//Example continue keyword
	//controlflow.ExampleContinue()

	//Example extract
	structure.ExampleStruct()

	//Message from main local
	fmt.Println("Welcome in learning world")

}
