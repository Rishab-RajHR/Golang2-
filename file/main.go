package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()

	content := "Hello World By Alex"
	_, errors := io.WriteString(file, content)
	if errors != nil {
		fmt.Println("Error while writing file: ", errors)
		return
	}

	fmt.Println("Successfully Created File")
}
