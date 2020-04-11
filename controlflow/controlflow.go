package controlflow

import "fmt"

func PrintHelloWordNtimes() {

	for i := 0; i < 10; i++ {
		fmt.Println("Hello Word! from for loop")
	}
}

func ExampleContinue() {

	i := 1
	for {

		i++
		//fmt.Println(i)
		if i > 10 {
			break
		}

		if i%2 != 0 {

			continue

		}

		fmt.Println(i)

	}
}
