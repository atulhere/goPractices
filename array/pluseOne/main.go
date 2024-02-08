/*
You are given a large integer represented as
an integer array digits, where each digits[i]
is the ith digit of the integer. The digits are
ordered from most significant to least significant
in left-to-right order. The large integer does not
contain any leading 0's.
Increment the large integer by one and return
the resulting array of digits.
*/

package main

import (
	"fmt"
)

func increment(array []int) []int {

	n := len(array)
	carry := 1

	for i := n - 1; i >= 0; i-- {

		sum := array[i] + carry
		array[i] = sum % 10
		carry = sum / 10

		if carry == 0 {

			break

		}
	}
	if carry > 0 {
		array = append([]int{carry}, array...)
	}

	return array
}

func main() {

	array := []int{1, 9}

	fmt.Println("The given array is ", array)

	result := increment(array)

	fmt.Println("The incremented array is ", result)

}
