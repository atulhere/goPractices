package main

import (
	"fmt"
)

func main() {
	hello()
}
func hello() {

	//https://go.dev/play/p/iuK_MkYXp24
	ch := make(chan string)

	go func() {
		//writing the value into the channel(ch) in one goroutine
		ch <- "Hello Maa"
	}()

	// Reading the value from channel(ch) in main goroutine
	message := <-ch
	fmt.Println(message)

}
