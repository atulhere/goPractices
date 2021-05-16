package programs

import (
	"fmt"
)

func NumbersInAscOrder() {

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {

			fmt.Printf("%d ", j)

		}
		fmt.Println()
	}
}
