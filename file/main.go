package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("exmaple.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()
}
