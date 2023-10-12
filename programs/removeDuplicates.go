package main

import (
	"fmt"
)

func main() {

	var age = 10 // Declare and initialize values
	age = 20
	var name string = "Atul" // Declare and initialize values
	var arr = []int{3, 5, 2, 1, 3, 2, 5}
	var uniqueArray = make(map[int]bool)
	var tempArray = []int{}

	for i := 0; i < len(arr); i++ {

		fmt.Println("The value is", arr[i])

		if uniqueArray[arr[i]] != true {
			uniqueArray[arr[i]] = true
			tempArray = append(tempArray, arr[i])

		}

	}

	fmt.Println("Unique Array is ")
	fmt.Println(tempArray)
}
