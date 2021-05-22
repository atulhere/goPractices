package slices

import "fmt"

func PrintSlices() {

	slice := []int16{1, 2, 3, 4, 5}

	fmt.Println("The value of initial Slice is", slice)

	//Slice a Slice

	sliceFirst := slice[2:3]

	fmt.Println("The value of first Slice is from slice", sliceFirst)

}

func AppendToSlices() {

	//define a slice

	x := []string{"a", "e", "i"}

	y := []string{"o", "u", "x"}

	//xy := append(x, "d", "b", "c") correct

	//xy := append(x, "d", "b", "c",10) ca not append differet tatatype values

	xy := append(x, y...)

	fmt.Println(xy)

}

func UseOfMake() {

	slice := make([]int16, 5, 7)

	fmt.Println("The default value of Slices")

	slice = []int16{2, 5, 6, 7, 8}
	//slice[6] = 9 // This is ot possible index out of range problem
	sliceAppend := append(slice, 9)
	fmt.Println(sliceAppend)
}
