package main

import "fmt"

func main() {

	//Initilize arr1 and arr2
	arr1 := []int{1, 2, 0, 3, 4, 5}
	arr2 := []int{0, 2, 3, 9, 7}

	//Create a tempArray array as a Hash
	tempArray := make(map[int]int)

	//Create an array to store common elements from both arrays
	commonArray := []int{}

	//Loop through arr1 and push the elements to Hash
	for _, v := range arr1 {

		tempArray[v] = v

	}

	/*
		Loop through arr2 and check if
		Has(tempArray) already have some elments of arr2
	*/
	commonElements := false
	for _, v := range arr2 {

		if tempArray[v] == v {
			commonElements = true

			//push common elments into a commonArray
			commonArray = append(commonArray, v)
		}
	}

	if commonElements {
		fmt.Println("The given arrays have common elments.")
		fmt.Println("The commany array is ", commonArray)

	} else {

		fmt.Println("The given arrays does not have any common elments.")
	}
}
