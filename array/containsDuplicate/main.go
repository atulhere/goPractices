package main

import "fmt"

func containsDuplicate(nums []int) bool {

	temp := make(map[int]bool)
	isDuplicate := false

	for i := 0; i < len(nums); i++ {

		if temp[nums[i]] {
			isDuplicate = true

		} else {
			temp[nums[i]] = true
		}
	}

	return isDuplicate

}

func main() {

	array := []int{0, 1, 2, 3, 5, 0}
	isD := containsDuplicate(array)

	fmt.Println("Is Duplicate is equal to ", isD)

}
