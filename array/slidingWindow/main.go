package main

import "fmt"

// From a given array find out a subarray of k length, which have max sum

// Sum of all subarrays by Brute force Approach

func SumOfSubArraysBruteForce(arr []int, length int) []int {

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

	}

	return sum
}

func MaxSumOfAllSubarraysBruteForce(arr []int, length int) int {

	// total number of elements in array
	n := len(arr)

	maxSum := 0
	for i := 0; i < (n - length + 1); i++ {
		sum := 0
		for j := i; j < (i + length); j++ {

			sum = sum + arr[j]

		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum

}

func MaxSumUsingSliwdingWindow(arr []int, lenght int) int {
	// total number of elements in array
	n := len(arr)

	// find sum of first k elements of array
	previousSum := arr[0] + arr[1] + arr[2]

	maxSum := previousSum
	nextSum := 0

	for i := 1; i < (n - lenght + 1); i++ {

		nextSum = (previousSum - arr[i-1]) + arr[(i+lenght)-1]

		if nextSum > maxSum {
			maxSum = nextSum
		}

		previousSum = nextSum

	}
	return maxSum
}

func main() {

	// let's define an array
	var arr = []int{3, 1, 2, 7, 9, 0, 1, 3, 7}

	fmt.Println("The Given Array is ", arr)

	// Size of subarrays
	var length int = 3

	sumOfSubArrays := SumOfSubArraysBruteForce(arr, length)
	fmt.Println("Array of sum of all sub arrays", sumOfSubArrays)

	maxSumOfAllSubarrays := MaxSumOfAllSubarraysBruteForce(arr, length)
	fmt.Println("The maximum sum of all possible subarray is: ", maxSumOfAllSubarrays)

	// Maximum sum of all subarrays of lenth k , using sliding window

	maxSum := MaxSumUsingSliwdingWindow(arr, length)
	fmt.Println("Maximum Sum of all subarrays of lenth k is ", maxSum)

}
