package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func testLoginController(t *testing.T) {

	// Set Gin to Test mode
	gin.SetMode(gin.TestMode)

	// Create a test router with the same routes as your main router
	router := setupRouter()

	// Create a request to test the /login route
	loginRequest := LoginRequest{
		Name:     "John",
		Email:    "john@example.com",
		LastName: "Doe",
		Password: "secretpassword",
	}

	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request using the test router
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body contains a JWT token
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if token, ok := response["token"]; !ok || token == "" {
		t.Errorf("Expected a non-empty JWT token in the response, got: %v", response)
	}
}
