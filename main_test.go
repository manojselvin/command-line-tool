package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Test Case: 1
func TestFetchTodos(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		mockResponse := `[{"userId": 1, "id": 1, "title": "Test Todo 1", "completed": false},
						  {"userId": 1, "id": 2, "title": "Test Todo 2", "completed": true},
						  {"userId": 2, "id": 3, "title": "Test Todo 3", "completed": false}]`
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Call the fetchTodos function with the mock server URL
	todos, err := fetchTodos(mockServer.URL)
	if err != nil {
		t.Fatalf("fetchTodos returned an error: %v", err)
	}

	// Expected result
	expected := []Todo{
		{UserID: 1, ID: 1, Title: "Test Todo 1", Completed: false},
		{UserID: 1, ID: 2, Title: "Test Todo 2", Completed: true},
		{UserID: 2, ID: 3, Title: "Test Todo 3", Completed: false},
	}

	// Check if the returned todos match the expected todos
	if !reflect.DeepEqual(todos, expected) {
		t.Errorf("fetchTodos returned unexpected todos. Got: %v, Expected: %v", todos, expected)
	}
}

// Test Case: 2
func TestFilterEvenTodos(t *testing.T) {
	todos := []Todo{
		{UserID: 1, ID: 1, Title: "Test Todo 1", Completed: false},
		{UserID: 1, ID: 2, Title: "Test Todo 2", Completed: true},
		{UserID: 2, ID: 3, Title: "Test Todo 3", Completed: false},
	}

	// Expected result after filtering even IDs
	expected := []Todo{
		{UserID: 1, ID: 2, Title: "Test Todo 2", Completed: true},
	}

	evenTodos := filterEvenTodos(todos)

	// Check if the returned todos match the expected todos
	if !reflect.DeepEqual(evenTodos, expected) {
		t.Errorf("filterEvenTodos returned unexpected todos. Got: %v, Expected: %v", evenTodos, expected)
	}
}
