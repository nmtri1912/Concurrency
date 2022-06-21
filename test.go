package main

import (
	"fmt"
	"time"
)

func printSth(comChan chan string) {
	i := 0

	for {
		i += 1
		select {
		case value := <-comChan:
			fmt.Println("Value", value)
			return
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	comChan := make(chan string)
	go printSth(comChan)
	time.Sleep(1 * time.Second)
	comChan <- "Finish"
}
