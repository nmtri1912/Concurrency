package main

import (
	"fmt"
	"log"
)

func printSth(msg string) chan string {
	// b2: waiter make card
	result := make(chan string)

	// b4: make food or drinks
	go func() {
		for i := 0; i < 5; i++ {
			result <- fmt.Sprintf("%s %s", msg, i)
		}
	}()

	// b3: waiter provide card for customer
	return result
}

func main() {
	// b1: customer order
	brigde := printSth("Hello")

	// b5: customer received
	for i := 0; i < 5; i++ {
		log.Println(<-brigde)
	}
}
