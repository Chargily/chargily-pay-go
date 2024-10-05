package chargily

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// struct represents the request sender
type RequestSender struct {
	hc *http.Client
	apiKey     string
}


// available functions to use
type RequestSenderI interface {
	SendRequest(method, endpoint string, body interface{}, result interface{}) error
}


//create a new request sender 
func NewRequestSender(apiKey string) RequestSenderI {
    return &RequestSender{
        hc: 			&http.Client{},
		apiKey: 		apiKey,
    }
}




// sendRequest sends an HTTP request and decodes the JSON response into the provided result interface.
func (rs * RequestSender) SendRequest(method, endpoint string, body interface{}, result interface{}) error {
	// Create a context with a 3-second timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var req *http.Request
	var err error

	// If the body is not nil, encode it as JSON
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}

		// Create request with body
		req, err = http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(jsonBody))
		if err!= nil {
            return fmt.Errorf("failed to create request with body: %v", err)
        }
	} else {
		// Create request without body
		req, err = http.NewRequestWithContext(ctx, method, endpoint, nil)
	}

	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers for the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ rs.apiKey)

	// Send the request using the provided HTTP client
	res, err := rs.hc.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()


	// Check if the status code is not in the 2xx range
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("failed with status: %s", res.Status)
	}


	// Decode the response body into the provided result interface (JSON)
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return nil
}
