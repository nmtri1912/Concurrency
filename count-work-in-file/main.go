package main

import (
	"log"
	"os"
	"strings"
)

func countFirstFile(firstChan chan int, filePath string, keyword string) {
	result := 0
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		firstChan <- result
		return
	}
	result = strings.Count(string(fileContent), keyword)
	firstChan <- result
	defer close(firstChan)
}

func countSecondFile(secondChan chan int, filePath string, keyword string) {
	result := 0
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		secondChan <- result
		return
	}
	result = strings.Count(string(fileContent), keyword)

	secondChan <- result
	defer close(secondChan)
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)

	go countFirstFile(firstChan, "1.txt", "hay")
	go countSecondFile(secondChan, "2.txt", "hay")

	log.Println("Total: ", <-secondChan+<-secondChan)
}
