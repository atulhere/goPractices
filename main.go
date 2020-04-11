package main

import (
	"fmt"

	"github.com/atulhere/goPractices/constant"
	"github.com/atulhere/goPractices/controlflow"
	"github.com/atulhere/goPractices/hello"
)

func main() {

	//Hello Word from Package
	hello.HellWord()

	//constant Practice
	constant.PlayWithConstant()

	//Example Control Flow
	controlflow.PrintHelloWordNtimes()

	//Example continue keyword
	controlflow.ExampleContinue()

	//Message from main local
	fmt.Println("Welcome in learning world")

}
