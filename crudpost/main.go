package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Alex Pandian",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	// Convert json data to string
	jsonReader := strings.NewReader(string(jsonData))

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}

	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response : ", string(data))

	fmt.Println("Response status : ", res.Status)
}

func main() {
	fmt.Println("Learning CRUD...")
	// performGetRequest()
	performPostRequest()
}
