package main

import "fmt"

func main() {

	array := []int{4, 2, -3, 1, 6}

	//found := zeroSumBruteForce(array)

	found := zeroSumUsingHash(array)

	if found {

		fmt.Println("Zero sum subarray found")
	} else {

		fmt.Println("Zero sum subarray not found")
	}

}

func zeroSumUsingHash(array []int) bool {

	temp := make(map[int]int)
	sum := 0
	isZeroSum := false

	for i := 0; i < len(array); i++ {

		sum = sum + array[i]

		if _, ok := temp[sum]; ok || (sum == 0) {

			isZeroSum = true

			fmt.Println("Temp Hasmap is ", temp)

			return isZeroSum

		}
		temp[sum] = sum

	}

	return isZeroSum

}

func zeroSumBruteForce(array []int) bool {

	isZeroSum := false
	for i := 0; i < len(array); i++ {

		sum := 0
		for j := i; j < len(array); j++ {

			sum = sum + array[j]

			if sum == 0 {

				isZeroSum = true

				return isZeroSum

			}

		}
	}

	return isZeroSum

}
