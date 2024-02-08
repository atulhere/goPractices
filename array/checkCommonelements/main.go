// Intersection of the two arrays
// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays
// and you may return the result in any order.

package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {

	tempArray1 := make(map[int]int)
	tempArray2 := make(map[int]int)

	var intersectArray []int

	//iterate nums1 array and store data in tempArray1
	for _, v := range nums1 {

		tempArray1[v]++
	}
	for _, v := range nums2 {

		tempArray2[v]++
	}

	for count1, v := range tempArray1 {

		// check if value of nums1 array is set in tempArray2
		if count2, ok := tempArray2[v]; ok {

			num := min(count1, count2)

			for i := 0; i < num; i++ {

				intersectArray = append(intersectArray, v)
			}
		}

	}
	return intersectArray

}
func main() {

	nums1 := []int{1, 2, 3, 4, 5, 1, 2}
	nums2 := []int{1, 2, 8, 9, 1, 1, 2}

	intersectArray := intersect(nums1, nums2)

	fmt.Println("The intersection of arrays is ", intersectArray)

}
