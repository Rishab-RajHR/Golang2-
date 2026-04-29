package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response ", err)
		return
	}
}
