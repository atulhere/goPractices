package main

import "fmt"

/* Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func TwoSumUsingBruteForce(arr []int, target int) []int {

	n := len(arr)

	var indexArray []int

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {

			if i != j {
				if arr[i]+arr[j] == target {

					indexArray = append(indexArray, i, j)

					return indexArray
				}
			}
		}
	}

	return indexArray

}

func twoSum(arr []int, target int) []int {

	n := len(arr)
	var tempMap = make(map[int]int)
	var lastIndex int
	var num int
	for i := 0; i < n; i++ {

		num = target - arr[i]

		if _, ok := tempMap[num]; ok {

			lastIndex = tempMap[num]

			return []int{lastIndex, i}
		}
		tempMap[arr[i]] = i

	}
	return []int{}
}

func main() {

	arr := []int{0, 1, 3, 5, 6, 9, 2, 8}

	// suppose target is 11
	indexArray := TwoSumUsingBruteForce(arr, 8)
	fmt.Println("The index of the array elements: ", indexArray)

	// suppose target is 4
	indexArr := twoSum(arr, 8)
	fmt.Println("The index of the array elements: ", indexArr)

}
