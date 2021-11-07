package main

import "fmt"

func main() {

	PrintMessageFromChannel()
}
func PrintMessageFromChannel() {

	ch := make(chan string)

	ch <- "Hello Maa"

	message := <-ch

	fmt.Println("This message is using from channel", message)

}
