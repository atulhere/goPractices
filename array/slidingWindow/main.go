package main

import "fmt"

// From a given array find out a subarray of k length, which have max sum

// Brute force Approach

func SumOfSubArraysByBruteForce(arr []int, length int) []int {

	// total number of elements in array
	n := len(arr)

	// let's declare a slice which will holds sum of all possible subarrays elements
	var sum []int

	for i := 0; i <= n-length; i++ {

		// temporary variable to hold the sum of three consicutive subarray elments
		var temp int
		for j := i; j < length+i; j++ {
			temp = temp + arr[j]
		}

		sum = append(sum, temp)

		fmt.Printf("Print sum of subarrays elements starting from the index %d  is %d ", i, sum)

	}

	return sum
}

func main() {

	// let's define an array
	var arr = []int{3, 1, 2, 7, 9, 0, 1, 3, 7}

	fmt.Println("The Given Array is ", arr)

	// Size of subarrays
	var length int = 3

	sumOfSubArrays := SumOfSubArraysByBruteForce(arr, length)

	fmt.Println("Subarray with Maximum sum is ", sumOfSubArrays)

}
