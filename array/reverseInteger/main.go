package main

import (
	"fmt"
	"math"
)

// Program to reverse a integer number

func reverse(n int) int {

	result := 0
	var remainder int

	k := 1

	// check if number is negative
	if n < 0 {

		k = -1
		n = n * k
	}

	for n > 0 {

		remainder = n % 10

		result = result*10 + remainder

		n = n / 10

	}

	limit := math.MaxInt32
	if result > limit || result < -limit {
		fmt.Println("The number is bigger than the Int32 limit")
		return 0
	}
	return result * k

}

func main() {

	var n int

	fmt.Println("Enter the integer to reverse")
	fmt.Scanf("%d", &n)

	result := reverse(n)

	fmt.Printf("The reverse of the Integer %d is %d", n, result)

}
