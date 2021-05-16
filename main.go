package main

import (
	//"goPractices/api"
	"goPractices/constant"
	"goPractices/hello"
	"goPractices/structure"

	//"goPractices/variable"
	"goPractices/programs"
	"goPractices/programs/array"
	"goPractices/programs/string"
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
	//fmt.Println("Welcome in learning world")

	//Example Pass by value
	//variable.PassByValue()

	//How to make api sample example
	//api.TestApis()

	//Print Number Patternn in Asc order

	programs.NumbersInAscOrder()

	//Program to practices arrays

	array.PrintArray()

	//Program to practice string
	string.PrintStringElement()

	// Program to swap two nummbers

	programs.SwapTwoNumbers(10, 20)

}
