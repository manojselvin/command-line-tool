package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Todo represents a TODO item
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Fetch TODO items from an API endpoint
	todos, err := fetchTodos("https://jsonplaceholder.typicode.com/todos?limit=40")
	if err != nil {
		fmt.Println("Error fetching TODOs:", err)
		return
	}

	// Filter even numbered TODOs
	evenTodos := filterEvenTodos(todos)

	// Output title and completion status of each TODO
	for _, todo := range evenTodos {
		fmt.Printf("Id: %d | Title: %s | Completed: %t\n", todo.ID, todo.Title, todo.Completed)
	}
}

// Fetches TODO items from the specified URL
func fetchTodos(url string) ([]Todo, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var todos []Todo
	err = json.NewDecoder(response.Body).Decode(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Filters even numbered TODOs
func filterEvenTodos(todos []Todo) []Todo {
	var evenTodos []Todo
	for _, todo := range todos {
		if todo.ID%2 == 0 {
			evenTodos = append(evenTodos, todo)
		}
		print(todo.ID)
		if len(evenTodos) == 20 {
			break
		}
	}
	return evenTodos
}
