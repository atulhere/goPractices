package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {

	//create object of pserson type
	p1 := person{Name: "Ram", Age: 90}

	var p2 person

	//marshal data from Object to JSON
	//first data will be converted into byte code (Array of byte)
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data in the byte code", data)
	//convert json string from the byte code
	jsonString := string(data)
	fmt.Println(jsonString)

	// now lets understand concept of unmarshaling using slightly different string values

	//st := `{"Name":"Ram","Age":90}`

	//now covert json string to byte code

	//byteCode := []byte(st)

	//now using byte code Unmarshal into Object
	json.Unmarshal(data, &p2)

	fmt.Println(p2.Name)

}
