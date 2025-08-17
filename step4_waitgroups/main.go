package main

import (
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Worker", id, "is working")
	time.Sleep(time.Duration(id) * time.Second)
	println("Worker", id, "finished")

}

func main() {

	var wg sync.WaitGroup
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	println("Main: All workers started, waiting for them to finish")
	wg.Wait()
	println("Main: All workers finished")
}
