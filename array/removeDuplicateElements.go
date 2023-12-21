package array

import "fmt"

func RemoveDuplicateElement(arr [5]int) {

	arrMap := make(map[int]int)
	var uniqueArray []int
	fmt.Println(arr)
	for _, v := range arr {
		if arrMap[v] != v {
			arrMap[v] = v
			uniqueArray = append(uniqueArray, v)
		}
	}

	fmt.Println("The Unique array is", uniqueArray)

}
