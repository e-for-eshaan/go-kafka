package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-server/api"
	"web-server/database"
)

func TestGetUsers(t *testing.T) {
	// Create a request to the "/users" endpoint
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function (getUsers) directly to serve the request
	api.GetUsers(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v", http.StatusOK, rr.Code)
	}

	// Parse the response body and check the content
	var users []api.User
	if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Perform additional checks on the users data as needed
	// For example, check the length of the users slice, verify specific data, etc.
	// ...
}

func TestGetProducts(t *testing.T) {
	// Create a request to the "/products" endpoint
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function (getProducts) directly to serve the request
	api.GetProducts(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v", http.StatusOK, rr.Code)
	}

	// Parse the response body and check the content
	var products []api.Product
	if err := json.Unmarshal(rr.Body.Bytes(), &products); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Perform additional checks on the products data as needed
	// For example, check the length of the products slice, verify specific data, etc.
	// ...
}

func TestCreateUser(t *testing.T) {
	// Define a new user payload in JSON format
	newUser := api.User{
		Name:   "Test User",
		Mobile: "1234567890",
	}

	// Convert the user payload to JSON
	payload, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user payload: %v", err)
	}

	// Create a request to the "/users" endpoint with POST method and the user payload
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}

	// Set the Content-Type header to indicate JSON data
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function (createUser) directly to serve the request
	api.CreateUser(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v", http.StatusCreated, rr.Code)
	}

	// Parse the response body and check the content (if needed)
	// ...

	// Perform additional checks on the database to verify the user was created (if needed)
	// ...
}

// You can add more test functions for other server endpoints or edge cases as needed.

func TestMain(m *testing.M) {
	// Call flag.Parse() here if TestMain uses flags

	// Initialize the database connection before running tests
	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.DB.Close()

	// Run the tests
	m.Run()
}
