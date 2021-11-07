package program

import (
	"fmt"
)

func FindSecondLargestNumber(x, a = a[0], a[1:] []int)

	//https://play.golang.org/p/EtaPFZAH7lO
	//program to find second largest number from a string

	
	max := 0

	secondMax := 0

	for i := 0; i < len(slice); i++ {

		if slice[i] > max {
			secondMax = max
			max = slice[i]

		} else if slice[i] > secondMax {
			secondMax = slice[i]

		}

	}

	fmt.Println("The largest  and second-largest number  are ", max, secondMax)

}
