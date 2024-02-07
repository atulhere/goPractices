package main

import "fmt"

func rotate(array []int, k int) []int {

	//length of an Array
	n := len(array)
	// Use a temporary Array
	// to store values after rotation
	var tempArray = make([]int, n)
	var tempIndex int

	for i := 0; i < n; i++ {

		tempIndex = (i + k) % n
		fmt.Println(tempIndex)
		tempArray[tempIndex] = array[i]

	}
	return tempArray

}

func main() {

	// Given Array
	array := []int{1, 2, 3, 4, 5, 6, 7}
	// rotate the array by k steps
	k := 3

	rArray := rotate(array, k)

	fmt.Println("Rotaed array is ", rArray)

}
