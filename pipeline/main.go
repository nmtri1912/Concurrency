package main

import "log"

func firstInput(nums ...int) chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < len(nums); i++ {
			c <- nums[i]
		}
		close(c)
	}()
	return c
}

func secondInput(fromFirst chan int) chan int {
	result := make(chan int)

	go func() {
		for item := range fromFirst {
			result <- item * 2
		}
		close(result)
	}()

	return result
}

func main() {
	firstChan := firstInput(1, 2, 3, 4, 5)
	secondChan := secondInput(firstChan)
	for item := range secondChan {
		log.Println("Received: ", item)
	}
}
