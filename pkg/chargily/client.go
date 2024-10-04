package chargily

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client is the structure that holds the API key , the endpoint for the Chargily API and the development mode.
type Client struct {
    apiKey   		string
    endpoint 		string
    client   		*http.Client // Allows for custom HTTP client configurations if needed
	mode            Mode
}



// NewClient initializes and returns a new Client with the given API key and endpoint.
func NewClient(apiKey , mode string) (*Client, error) {

    //Set the API base URL based on the mode provided
    var api_baseUrl string

    //Verify the mode parameter
    // Verify the mode parameter
    if mode == "prod" {
        api_baseUrl = ProdAPIBaseUrl
    } else if mode == "test" {
        api_baseUrl = TestAPIBaseUrl
    } else {
        return nil, ErrInvalidMode
    }

    fmt.Println(apiKey,"\n", api_baseUrl,"\n", mode)
    
    //return the client with it's configurations
    return &Client{
        apiKey:   	apiKey,
        endpoint: 	api_baseUrl,
        client:   	&http.Client{},
		mode:       Mode(mode), //test: for testing/development stage , prod: for production applications
    }, nil 
}



//retrieve the balance example 
func (c *Client) GetBalance() (*Balance, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Build the request
	url := c.endpoint + "balance"
    fmt.Println(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Add the authorization header
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Send the request
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()


    
    // Check if status code is not 200
    if res.StatusCode != http.StatusOK {
        // Try to decode the error message from the response body
        var apiError struct {
            Message string `json:"message"` 
        }

        if err := json.NewDecoder(res.Body).Decode(&apiError); err != nil {
            return nil, fmt.Errorf("failed to fetch balance with status: %s and could not decode error message: %v", res.Status, err)
        }

        // Return the error message from the response
        return nil, fmt.Errorf("failed to fetch balance with status: %s, error: %s", res.Status, apiError.Message)
    }



	// Parse the response body
	var balance Balance
	if err := json.NewDecoder(res.Body).Decode(&balance); err != nil {
		return nil, err
	}

	// Return the parsed balance object
	return &balance, nil
}