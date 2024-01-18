package main

import (
	"fmt"
	"goPractices/constant"
	"goPractices/function"
	"goPractices/hello"
	"goPractices/method"
	"goPractices/programs/slices"
	"goPractices/programs/string"
)

func main() {

	TestGreetings()

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

	//programs.NumbersInAscOrder()

	//Program to practice string
	string.PrintStringElement()

	// Program to swap two nummbers
	//programs.SwapTwoNumbers(10, 20)

	//Practice Slices
	slices.PrintSlices()

	slices.AppendToSlices()

	slices.UseOfMake()

	//Example Functionn
	function.SumOfNummbers(5, 4)

	//Example Variadic params
	function.VariadicParameters(5, 5, 3, 2, 8)

	//Unfurling Slice
	s := []int{1, 2, 3, 4, 5, 5, 5}
	function.UnfurlingSlice(s...)

	//Exmaple of anonymous function
	function.ExampleAnonymousFunction()

	//Example Method
	method.ExmapleMethod()

	//Find factorial using recursion
	number := function.Factorial(3)
	fmt.Println("The factorial of 3  is ", number)

	//Find factorial using recursion
	n := function.FactorialWithoutRecursion(3)
	fmt.Println("The factorial of 3  is ", n)

}
