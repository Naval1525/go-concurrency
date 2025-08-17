package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents work to be done
type Job struct {
	ID   int
	Name string
	Size int // Simulates different job sizes
}

// Result represents the output of completed work
type Result struct {
	Job    Job
	Worker int
	Output string
	Time   time.Duration
}

// Worker processes jobs
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("🚀 Worker %d started\n", id)

	for job := range jobs { // Keep receiving until channel is closed
		start := time.Now()

		fmt.Printf("👷 Worker %d processing job %d (%s)\n", id, job.ID, job.Name)

		// Simulate work (different jobs take different time)
		workTime := time.Duration(job.Size) * 100 * time.Millisecond
		time.Sleep(workTime)

		// Create result
		result := Result{
			Job:    job,
			Worker: id,
			Output: fmt.Sprintf("Job %d completed by worker %d", job.ID, id),
			Time:   time.Since(start),
		}

		results <- result
		fmt.Printf("✅ Worker %d finished job %d\n", id, job.ID)
	}

	fmt.Printf("🛑 Worker %d shutting down\n", id)
}

func main() {
	fmt.Println("=== STEP 8: WORKER POOL PATTERN ===")
	fmt.Println("Combining: Goroutines + Channels + WaitGroup + Buffering")

	const numWorkers = 3
	const numJobs = 12

	// Create channels
	jobs := make(chan Job, numJobs)       // Buffered jobs channel
	results := make(chan Result, numJobs) // Buffered results channel

	// Start workers
	var wg sync.WaitGroup
	fmt.Printf("\n🏭 Starting %d workers...\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	fmt.Printf("\n📋 Sending %d jobs...\n", numJobs)
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Name: fmt.Sprintf("task-%d", i),
			Size: rand.Intn(5) + 1, // Random job size 1-5
		}
		jobs <- job
	}
	close(jobs) // No more jobs - workers will exit when done

	// Close results channel when all workers are done
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Then close results channel
	}()

	// Collect and display results
	fmt.Println("\n📊 Results:")
	var totalTime time.Duration
	var processedJobs int

	for result := range results {
		fmt.Printf("📦 %s (took %v)\n", result.Output, result.Time)
		totalTime += result.Time
		processedJobs++
	}

	fmt.Printf("\n🎯 SUMMARY:\n")
	fmt.Printf("• Jobs processed: %d\n", processedJobs)
	fmt.Printf("• Total processing time: %v\n", totalTime)
	fmt.Printf("• Average time per job: %v\n", totalTime/time.Duration(processedJobs))
	fmt.Printf("• Workers used: %d\n", numWorkers)

	fmt.Println("\n🎓 CONGRATULATIONS!")
	fmt.Println("You've learned all key Go concurrency concepts:")
	fmt.Println("✅ Goroutines - Concurrent execution")
	fmt.Println("✅ Channels - Communication between goroutines")
	fmt.Println("✅ Select - Handle multiple channels")
	fmt.Println("✅ WaitGroups - Synchronization")
	fmt.Println("✅ Mutex - Protect shared data")
	fmt.Println("✅ Buffered Channels - Performance optimization")
	fmt.Println("✅ Worker Pool - Real-world production pattern")

	fmt.Println("\n🚀 You're now ready to build concurrent applications in Go!")
}
