package main

func main() {
	chanInt := make(chan int)
	chanInt <- 5
	// <- chanInt
}
