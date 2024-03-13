package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	counter++
}

func decrement() {
	counter++
}

func main() {

	go increment()
	go decrement()

	time.Sleep(time.Second)

	fmt.Println("Value of the counter is : ", counter)

	// go run --race race/main.go
	//WARNING: DATA RACE

}
