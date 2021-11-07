package array

import "fmt"

//Remove Duplicate Elements from arrary
func RemoveDuplicateElement() {

	array := []int{1, 2, 3, 4, 5, 1, 2, 3}
	fmt.Println("The array is ", array)

	uniqueArray := []int{}

	for i := 0; i < len(array); i++ {

		//check if the element in array

		duplicate := CheckIfElementExist(array[i], uniqueArray)

		if !duplicate {

			uniqueArray = append(uniqueArray, array[i])

		}

	}

	fmt.Println(" Final Unique Array is like", uniqueArray)

}

func CheckIfElementExist(x int, y []int) bool {

	for j := 0; j < len(y); j++ {

		if x == y[j] {

			return true
		}
	}
	return false
}
