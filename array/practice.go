package main

import "fmt"

func main() {

	//Declare an array

	arr := []int{3, 5, 6}

	for k, v := range arr {

		if k == 2 {
			arr[2] = 10

			fmt.Printf("With Key %d value is %d \n", k, arr[2])
			fmt.Printf("With Key %d value is %d \n", k, v)
		}

	}

}
