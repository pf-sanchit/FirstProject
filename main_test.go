package main

import (
	"bytes"
	"context"
	"encoding/json"
	"firstProject/service/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloEndpoint ensures that the Hello endpoint responds as expected.
func TestHelloEndpoint(t *testing.T) {
	// Create a mock service
	service := v1.NewService()
	endpoints := v1.Endpoints{
		HelloEndpoint: v1.MakeHelloEndpoint(service),
	}

	// Create a context and handler
	ctx := context.Background()
	handler := v1.MyHTTPServer(ctx, endpoints)

	// Create the request payload (JSON body)
	requestBody := map[string]string{
		"name": "myName",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a new HTTP request with the JSON body
	req, err := http.NewRequest("GET", "/hello", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the HTTP response
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	// Check if the status code is OK
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Expected status 200 OK, but got %v", status)
	}

	// Check if the response body is the correct JSON
	var response map[string]string
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Validate the response JSON
	expectedMessage := "Hello myName"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message '%s', but got '%s'", expectedMessage, response["message"])
	}
}
