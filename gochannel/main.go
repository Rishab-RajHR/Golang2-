package main

import (
	"fmt"
	"sync"
)

func worker(i int) {
	fmt.Printf("Worker %d started\n", i)
	// some task is happening
	fmt.Printf("Worker %d end\n", i)
}

func main() {
	fmt.Println("Explore goroutine started")

	var wq sync.WaitGroup
	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wq.Add(1) // Increment the WaitGroup counter
		go worker(i)
	}
	fmt.Println("Workers task complete")
}
