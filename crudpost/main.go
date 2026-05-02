package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response: ", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading : ", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todos Todo
	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		fmt.Println("Error decoding : ", err)
		return
	}
	fmt.Println("Todo: ", todos)
	fmt.Println("Title response: ", todos)
	fmt.Println("completed response: ", todos)
}

func main() {
	fmt.Println("Learning CRUD...")
	performGetRequest()

}
