package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "http://localhost:5173/"
	fmt.Printf("Type of URL: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Can't parse URL ", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)
}
