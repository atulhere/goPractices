package main

import "fmt"

func main() {

	//Define map using map literal
	s := map[string]string{"word": "maa", "village": "Ahamadpur"}

	fmt.Println("Print map details", s)

	//loop through map and print key and values
	for k, v := range s {
		fmt.Printf("%s: %s \n", k, v)
	}

	//delete element from map using delete
	delete(s, "village")
	fmt.Println("Now map is", s)
}
