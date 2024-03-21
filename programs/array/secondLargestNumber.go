package array

func SecondLargest(arr []int) int {

	var largest, secondLargest int

	for i := 0; i < len(arr); i++ {

		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {

			secondLargest = arr[i]
		}

	}

	return secondLargest

}
