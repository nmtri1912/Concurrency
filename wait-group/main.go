package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	var result int
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
		result += i
	}
	defer wg.Done()
}

func printChar(wg *sync.WaitGroup) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c ", i)
		result += string(i)
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumber(&wg)
	go printChar(&wg)

	wg.Wait()
}
