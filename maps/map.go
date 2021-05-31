package maps

import "fmt"

func ExampleMap() {

	s := map[string]string{"word": "maa", "village": "Ahamadpur"}

	fmt.Println("The map value is", s)
}

func Practice() {

	s := map[string]string{"name": "maa", "village": "Ahamadpur", "post": "Ahamadpur", "district": "Jaunpur", "state": "UP", "country": "India"}

	fmt.Println("Print Address")
	for k, v := range s {

		fmt.Printf("%s: %s \n", k, v)
	}

}

func DeleteFromMap() {

	//delete()
}
