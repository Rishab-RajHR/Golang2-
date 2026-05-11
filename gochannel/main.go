package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", i)
	// some task is happening
	fmt.Printf("Worker %d end\n", i)
}

func main() {
	fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	// wait for all workers to finish
	wg.Wait()
	fmt.Println("Workers task complete")
}
