package main

import "time"

func worker(name string, done chan bool) {
	println(name)
	time.Sleep(2 * time.Second)
	println("finished")
	done <- true
}

func workerWithMessage(name string, result chan string) {
	println(name)
	time.Sleep(1 * time.Second)
	result <- "done"
}

func main() {
	// using done channel
	done := make(chan bool)
	go worker("Naval", done)
	println("waiting to finish task")
	<-done //wait it will block the code
	println("task finished")

	//recieving data
	result := make(chan string)
	go workerWithMessage("Devanshi", result)
	message := <-result
	println("received message:", message)
}
