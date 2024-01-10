package main

import "fmt"

func main() {

	//Declare an array
	var exampleArray [3]int

	//Initilize the value to array
	exampleArray[0] = 3

	fmt.Println("The value of the array is ", exampleArray)

	//Print the type of array
	var arr [3]int
	fmt.Printf("The type of array is %T\n", arr)

	//Print the type of slice
	var slice []int
	fmt.Printf("The type of array is %T\n", slice)

	//Initilize some values to a slice

	var sli []int = []int{1, 0, 2, 3, 4, 5}

	//Print len and capicity of slice

	fmt.Printf("The len of sli is %d and cap is %d \n ", len(sli), cap(sli))
	s := sli[2:5]
	fmt.Println("The value of the sli is ", s)
	fmt.Printf("The len of s is %d andcap is %d \n ", len(s), cap(s))
}
