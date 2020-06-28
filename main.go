package main

import (
	"fmt"

	"github.com/atulhere/goPractices/constant"
	"github.com/atulhere/goPractices/hello"
	strucure "github.com/atulhere/goPractices/structure"
)

func main() {

	//Hello Word from Package
	hello.HellWord()

	//constant Practice
	constant.PlayWithConstant()

	//Example Control Flow
	//controlflow.PrintHelloWordNtimes()

	//Example continue keyword
	//controlflow.ExampleContinue()

	//Example extract
	strucure.ExampleStruct()

	//Message from main local
	fmt.Println("Welcome in learning world")

}
