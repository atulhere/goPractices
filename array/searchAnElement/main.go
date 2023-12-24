package main

import "fmt"

func main() {

	var arr [5]int
	fmt.Println("Please enter integer array up to 5 elements")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("The array is", arr)

	fmt.Println("Pleas enter number to be searched!")

	var searchValue int
	fmt.Scanf("%d", &searchValue)

	flag := false

	for i := 0; i < 5; i++ {
		if arr[i] == searchValue {
			flag = true
			fmt.Println("The number found at index", i)
		}
	}
	if !flag {
		fmt.Printf("The value is not found in the given array")
	}

}
