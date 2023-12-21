// find of sum of all even number in the array
package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("The value of default array is", a)
	sum := 0
	for _, v := range a {

		if v%2 == 0 {
			sum = sum + v
		}
	}

	fmt.Println("The sum of all even numbers in given array is ", sum)

}
