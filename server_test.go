package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler_ValidCredentials(t *testing.T) {
	// Create a new request with valid username and password
	req, err := http.NewRequest("POST", "/login?username=admin&password=password", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the login handler
	loginHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Assert the response body
	expectedResponse := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.RT0gX6rcwszWZUGIAqTAy8VHvxic_35LBT8Mh1zocDM"}`
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("expected response body '%s' but got '%s'", expectedResponse, rr.Body.String())
	}
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	// Create a new request with invalid username and password
	req, err := http.NewRequest("POST", "/login?username=admin&password=wrongpassword", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the login handler
	loginHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d but got %d", http.StatusUnauthorized, rr.Code)
	}

	// Assert the response body
	expectedResponse := "Invalid username or password"
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("expected response body '%s' but got '%s'", expectedResponse, rr.Body.String())
	}
}
