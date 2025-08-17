package main

import (
	"fmt"
	"time"
)

func fastWorker(result chan string) {
	time.Sleep(1 * time.Second)
	result <- "fast worker done"
}

func slowWorker(result chan string) {
	time.Sleep(3 * time.Second)
	result <- "slow worker done"
}

func main() {
	//select
	ch1 := make(chan string)
	ch2 := make(chan string)
	go fastWorker(ch1)
	go slowWorker(ch2)
	//select which ever comes first
	select {
	case msg := <-ch1:
		println("Received from fast worker:", msg)
	case msg := <-ch2:
		println("Received from slow worker:", msg)
	}
	//with timeout
	slowchn := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		slowchn <- "slow worker done"
	}()
	select {
	case msg := <-slowchn:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Operation took too long")
	}

	//non blocking
	data := make(chan string)
	select {
	case msg := <-data:
		fmt.Println("Received:", msg)
	default:
		println("No data")
	}
}
