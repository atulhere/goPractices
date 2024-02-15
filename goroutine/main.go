package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	url := []string{
		"https://google.com",
		"https://linkedin.com",
		"https://tuscanyyy.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	async := true //we can change this value to test sync/async behaviour
	callUrls(url, async)

}

func callUrls(url []string, async bool) {
	start := time.Now()
	fmt.Println("Process Started At ", start)
	if !async {
		for _, v := range url {
			syncCall(v)
		}
	} else {
		for _, v := range url {
			go asyncCall(v)
			wg.Add(1)
		}
		wg.Wait()

	}

	end := time.Since(start)
	fmt.Println("Total time taken in the process is ", end)

}

func asyncCall(url string) {

	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("There is an error while hitting %s\n", url)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Output while hitting %s is %d \n", url, response.StatusCode)
	}
}

func syncCall(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("There is an error while hitting %s\n", url)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Output while hitting %s is %d \n", url, response.StatusCode)
	}
}

// Basic program goroutine
//https://go.dev/play/p/a2PWPiFf69C
