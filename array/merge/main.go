package main

import "fmt"

func main() {

	array1 := []int{1, 2, 3}
	m := len(array1)
	array2 := []int{2, 3, 5, 6}
	n := len(array2)
	array := make([]int, m+n)

	i, j, k := 0, 0, 0

	fmt.Println("First sorted array: ", array1)
	fmt.Println("Second sorted array: ", array2)

	for (i < m) && (j < n) {

		if array1[i] <= array2[j] {

			array[k] = array1[i]
			i++
			k++
		} else {
			array[k] = array2[j]
			j++
			k++
		}

	}

	for i < m {

		array[k] = array1[i]
		i++
		k++

	}
	for j < n {
		array[k] = array2[j]
		j++
		k++
	}
	fmt.Println("The merged sorted array: ", array)
}
