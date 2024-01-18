package main

import "fmt"

func main() {

	// Given array
	arr := []int{0, 3, 2, 0, 0, 1, 2, 0, 5, 0, 0}

	// let's assume we need find all occurances of 2

	element := 9
	numOfOccurance := 0

	for _, v := range arr {
		if v == element {
			numOfOccurance++
		}
	}
	if numOfOccurance > 0 {
		fmt.Printf("The given element %d is repeated %d times. \n", element, numOfOccurance)
	} else {
		fmt.Println("The given element is not found in the array")
	}

}
