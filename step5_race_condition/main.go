package main

import "sync"

func main() {
	//race condition
	var counter int //shared
	var wg sync.WaitGroup
	numGoRoutines := 5
	increment := 1000
	println("Expected final result: %d\n", numGoRoutines*increment)
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < increment; j++ {
				counter++
			}
			println(id)

		}(i)

	}
	wg.Wait()
	println("Final result: ", counter)
	println("Expected final result:", numGoRoutines*increment)
}
