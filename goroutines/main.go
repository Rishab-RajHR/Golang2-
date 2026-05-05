package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(100 * time.Millisecond) // Simulating some work
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi Alex :)")
}

func main() {
	fmt.Println("Learning goroutines")
	sayHello()
	sayHi()
}
