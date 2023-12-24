package main

import "fmt"

func main() {

	arrMap := make(map[int]int)
	var uniqueArray []int
	arr := []int{1, 2, 1, 3, 2, 1}
	for _, v := range arr {
		if arrMap[v] != v {
			arrMap[v] = v
			uniqueArray = append(uniqueArray, v)
		}
	}

	fmt.Println("The Unique array is", uniqueArray)

}
