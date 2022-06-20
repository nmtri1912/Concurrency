package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func getURL(commuChan chan string, url string) {
	defer close(commuChan)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	commuChan <- string(content)
	defer resp.Body.Close()
}

func main() {
	firstChan := make(chan string)
	secondChan := make(chan string)
	thirdChan := make(chan string)

	f, err := os.OpenFile("./save.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return
	}

	listChan := []chan string{firstChan, secondChan, thirdChan}
	listUrl := []string{"https://youtube.com", "https://google.com", "https://bring.com"}

	for i := 0; i < len(listUrl); i++ {
		go getURL(listChan[i], listUrl[i])
	}

	for i := 0; i < len(listUrl); i++ {
		_, err := f.WriteString(<-listChan[i])
		if err != nil {
			return
		}
	}
}
