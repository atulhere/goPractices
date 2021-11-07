package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type response struct {
	status  bool
	message string
}

func main() {

	res := response{true, "All is Well"}
	fmt.Println(res)

	//create object of pserson type
	p1 := person{Name: "Atul", Age: 35}

	var p2 person

	//marshal data from Object to JSON

	//first data will be converted into byte code
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}

	//json string from byte code
	jsonString := string(data)
	fmt.Println(jsonString)

	// now lets understand concept of unmarshaling using slightly different string values

	st := `{"Name":"Ram","Age":90}`

	//now covert json string to byte code

	byteCode := []byte(st)

	//now using byte code Unmarshal into Object
	json.Unmarshal(byteCode, &p2)

	fmt.Println(p2.Name)

}
