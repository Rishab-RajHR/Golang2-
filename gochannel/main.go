package main

import "fmt"

func worker(i int) {
	fmt.Printf("Worker %d started", i)
	// some task is happening
	fmt.Printf("Worker %d end", i)
}

func main() {
	fmt.Println("Explore goroutine started")

	// Start 3 worker goroutines
	for i := 0; i < 3; i++ {
		worker(i)
	}
	fmt.Println("Workers task complete")
}
