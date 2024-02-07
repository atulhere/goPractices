// Intersection of the two arrays
package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	// Create frequency maps for each array
	freqMap1 := make(map[int]int)
	freqMap2 := make(map[int]int)

	for _, num := range nums1 {
		freqMap1[num]++
	}

	for _, num := range nums2 {
		freqMap2[num]++
	}

	// Find the intersection
	var result []int
	for num, count1 := range freqMap1 {
		if count2, ok := freqMap2[num]; ok {
			// Common element found, add to result the minimum frequency
			minFreq := min(count1, count2)
			for i := 0; i < minFreq; i++ {
				result = append(result, num)
			}
		}
	}

	return result
}

func main() {

	nums1 := []int{1, 2, 3, 0, 9, 2}

	nums2 := []int{2, 3, 4, 5, 8, 9, 2, 9, 8}

	array := intersect(nums1, nums2)

	fmt.Println("The intersection of the array is ", array)

}
