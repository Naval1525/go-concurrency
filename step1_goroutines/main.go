package main

import "time"

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		println("Hello", name)
		time.Sleep(1 * time.Second)
	}
	println("finished")
}

func main() {
	//sequential execution
	start := time.Now()
	sayHello("Naval")
	sayHello("Adi")
	elapsed := time.Since(start)
	println("Time taken for sequential execution:", elapsed.Seconds())
	start = time.Now()
	go sayHello("Devanshi")
	go sayHello("None")
	time.Sleep(4 * time.Second)
	elapsed = time.Since(start)
	println("Time taken for concurrent execution:", elapsed.Seconds())
}
