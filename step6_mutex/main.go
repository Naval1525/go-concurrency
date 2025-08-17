package main

import "sync"

type safeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *safeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func (c *safeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	//mutex
	counter := &safeCounter{}
	var wg sync.WaitGroup
	numGoRoutines := 5
	increment := 1000
	println("Expected final result: %d\n", numGoRoutines*increment)
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < increment; j++ {
				counter.Increment()
			}
			println(id)
		}(i)
	}
	wg.Wait()
	println("Final result: ", counter.Get())
	println("Expected final result:", numGoRoutines*increment)
}
