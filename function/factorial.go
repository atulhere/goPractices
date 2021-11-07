package function

//write a program for factorial using recursion

// Write another program for fatorial without recursion

func Factorial(n int) int {

	if n == 0 {
		return 1
	}
	number := n * Factorial(n-1)

	return number

}

func FactorialWithoutRecursion(n int) int {

	number := 1

	for i := n; i >= 1; i-- {

		number = i * number

	}
	return number

}
