package main

import (
	"fmt"
	"log"
)

func printNumber(numChan chan int) {
	var result int
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
		result += i
	}
	numChan <- result
}

func printChar(charChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c ", i)
		result += string(i)
	}
	charChan <- result
}

func main() {
	numChan := make(chan int)
	charChan := make(chan string)

	go printNumber(numChan)
	go printChar(charChan)

	log.Println("numChan: ", <-numChan)
	log.Println("charChan: ", <-charChan)
}
