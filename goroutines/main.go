package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(2000 * time.Millisecond) // Simulating some work
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi Alex :)")
}

func main() {
	fmt.Println("Learning goroutines")
	sayHello()
	sayHi()

	// Wait for a moment to allow the goroutine to finish
	time.Sleep(1000 * time.Millisecond)
}
