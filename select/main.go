package main

import (
	"fmt"
	"time"
)

func googleSearch(result chan string) {
	time.Sleep(3* time.Second)
	result <- "found from Google"
}

func bingSearch(result chan string) {
	time.Sleep(4*time.Second)
	result <- "found from Bing"
}

func main() {
	googleChan := make(chan string)
	bingChan := make(chan string)
	go googleSearch(googleChan)
	go bingSearch(bingChan)

	select {
	case result := <-googleChan:
		fmt.Println(result)
	case result := <-bingChan:
		fmt.Println(result)
	}
	fmt.Println("Finished")
}
