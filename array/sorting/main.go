package main

import "fmt"

// Program to sort a given array using bubble sort
func bubbleSort(arr []int) []int {

	length := len(arr)
	var swap bool
	for i := 0; i < length-1; i++ {
		swap = false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap = true
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp

			}
		}
		if !swap {
			break
		}

		fmt.Println("Print the index of i", i)
	}
	return arr
}

func main() {

	arr := []int{1, 0, 2, 7, 9, 11, 5, 8}

	fmt.Println("Unsorted Array: ", arr)
	//Sort Array using bubble sort
	sortedArray := bubbleSort(arr)
	fmt.Println("Sorted Array: ", sortedArray)

}
