package main

import (
	"fmt"
	"log"
)

func printSth(msg string) chan string {
	// b2: waiter make card
	c := make(chan string)

	// b4: make food or drinks
	go func() {
		for i := 0; i <= 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	// b3: waiter provide card for customer
	return c
}

func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}
		}
	}()
	return c
}
func main() {
	// b1: customer order
	coffee := printSth("coffee")
	bread := printSth("bread")

	serve := fanIn(coffee, bread)
	// b5: customer received
	for i := 0; i <= 5; i++ {
		log.Println(<-serve)
	}
}
