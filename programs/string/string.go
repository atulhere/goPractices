package string

import "fmt"

func PrintStringElement() {

	var string string

	string = "Maa"

	fmt.Println("The lenght of the string is ", len(string))

	fmt.Println("The Strings consist of")
	for i := 0; i < len(string); i++ {

		fmt.Printf("%c", string[i])
	}

}
