package unit_tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/stretchr/testify/assert"
)

func TestNewRequestSender(t *testing.T) {
	apiKey := "test_api_key"
	rs := chargily.NewRequestSender(apiKey)

	assert.NotNil(t, rs)
	assert.IsType(t, &chargily.RequestSender{}, rs)
	
	// Type assertion to access the unexported field
	//concreteRS, ok := rs.(*chargily.RequestSender)
	//assert.True(t, ok)
	//assert.Equal(t, apiKey, concreteRS.apiKey)
	//assert.NotNil(t, concreteRS.hc)
}

func TestSendRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request headers
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Bearer test_api_key", r.Header.Get("Authorization"))

		// Check request method
		assert.Equal(t, "POST", r.Method)

		// Decode the request body
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err)
		assert.Equal(t, "test_value", requestBody["test_key"])

		// Send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"response": "success"})
	}))
	defer server.Close()

	// Create RequestSender
	rs := chargily.NewRequestSender("test_api_key")

	// Prepare request body and result
	requestBody := map[string]string{"test_key": "test_value"}
	var result map[string]string

	// Send request
	err := rs.SendRequest("POST", server.URL, requestBody, &result)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "success", result["response"])
}

func TestSendRequestError(t *testing.T) {
	// Create a test server that always returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create RequestSender
	rs := chargily.NewRequestSender("test_api_key")

	// Send request
	var result map[string]string
	err := rs.SendRequest("GET", server.URL, nil, &result)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with status: 500")
}

func TestSendRequestTimeout(t *testing.T) {
	// Create a test server that sleeps longer than the timeout
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Second)
	}))
	defer server.Close()

	// Create RequestSender
	rs := chargily.NewRequestSender("test_api_key")

	// Send request
	var result map[string]string
	err := rs.SendRequest("GET", server.URL, nil, &result)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}