package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== STEP 7: BUFFERED CHANNELS ===")

	// Example 1: Unbuffered channel (blocks immediately)
	fmt.Println("\n1. Unbuffered channel:")
	unbuffered := make(chan string) // Capacity = 0

	go func() {
		fmt.Println("Goroutine: About to send to unbuffered channel...")
		unbuffered <- "Hello"
		fmt.Println("Goroutine: Sent to unbuffered channel!")
	}()

	time.Sleep(1 * time.Second) // Let goroutine try to send
	fmt.Println("Main: About to receive...")
	msg := <-unbuffered // This unblocks the sender
	fmt.Println("Main: Received:", msg)

	// Example 2: Buffered channel (doesn't block until full)
	fmt.Println("\n2. Buffered channel:")
	buffered := make(chan string, 3) // Capacity = 3

	// Send 3 items without any receivers (doesn't block!)
	buffered <- "Item 1"
	fmt.Println("Sent Item 1 (didn't block)")
	buffered <- "Item 2"
	fmt.Println("Sent Item 2 (didn't block)")
	buffered <- "Item 3"
	fmt.Println("Sent Item 3 (didn't block)")

	// Now channel is full, next send would block
	fmt.Println("Channel is now full (3/3)")

	// Receive them
	for i := 0; i < 3; i++ {
		item := <-buffered
		fmt.Println("Received:", item)
	}

	// Example 3: Producer-Consumer with buffer
	fmt.Println("\n3. Producer-Consumer pattern:")
	jobs := make(chan int, 5) // Buffer of 5 jobs

	// Producer (fast)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("Producing job %d\n", i)
			jobs <- i
			time.Sleep(100 * time.Millisecond) // Fast producer
		}
		close(jobs) // Signal no more jobs
	}()

	// Consumer (slower)
	time.Sleep(1 * time.Second) // Start consuming after a delay
	fmt.Println("Starting consumer...")

	for job := range jobs { // Receive until channel is closed
		fmt.Printf("Processing job %d\n", job)
		time.Sleep(300 * time.Millisecond) // Slow consumer
	}

	fmt.Println("\n✅ Key Learning: Buffered channels improve performance")
	fmt.Println("• Unbuffered: Sender blocks until receiver is ready")
	fmt.Println("• Buffered: Sender only blocks when buffer is full")
	fmt.Println("• Use buffering to handle different producer/consumer speeds")
	fmt.Println("❌ Problem: How to combine everything into real pattern?")
}
