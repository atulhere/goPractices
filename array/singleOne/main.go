/*
	Given a non-empty array of integers nums,

every element appears twice except for one.
Find that single one.
You must implement a solution with
a linear runtime complexity and use only
constant extra space.
*/
package main

import "fmt"

func singleNumber(nums []int) int {
	temp := make(map[int]bool)

	singleNumber := -1
	for i := 0; i < len(nums); i++ {
		if temp[nums[i]] {
			delete(temp, nums[i])

		} else {
			temp[nums[i]] = true
		}
	}

	for k, _ := range temp {
		singleNumber = k

	}
	return singleNumber

}

func main() {

	array := []int{8, 4, 5, 5, 8, 4}

	singleNumber := singleNumber(array)

	if singleNumber != -1 {
		fmt.Println("The single number is ", singleNumber)
	} else {
		fmt.Println("There is no such element in the array")
	}

}
