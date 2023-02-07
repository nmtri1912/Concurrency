package main

import (
	"fmt"
	"time"
)

func printSth(c chan int) {
	for {
		select {
		case <-c:
			fmt.Println("finish")
			return
		default:
			fmt.Println("Loop")
		}
	}
}

func main() {
	c := make(chan int)
	go printSth(c)

	time.Sleep(3 * time.Second)
	c <- 5
}

// dsadassda
