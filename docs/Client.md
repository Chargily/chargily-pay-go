# Chargily Client Documentation

## Overview

The `chargily` package provides a `Client` structure that facilitates interaction with the Chargily payment API. This client serves as a gateway to access various functionalities offered by the Chargily platform, including balance inquiries, customer management, product pricing, payment link generation, checkout processes, and webhook handling.

## Client Structure

### Definition

```go
type Client struct {
    apiKey   		string
    endpoint 		string
    rs              utils.RequestSenderI // RequestSender for custom HTTP requests
    mode            Mode
    Balance         *Balance              // Access balance-related functionalities
    Customers       *Customers            // Access customer management functionalities
    Prices          *Prices               // Access pricing-related functionalities
    Products        *Products             // Access product management functionalities
    PaymentLinks    *PaymentLinks         // Access payment link functionalities
    Checkouts       *Checkouts            // Access checkout functionalities
    Webhook         *Webhook              // Access webhook management functionalities
}
```

### Fields

- `apiKey`: A string representing the API key required to authenticate requests to the Chargily API.
- `endpoint`: A string representing the base URL of the Chargily API, which varies based on the operational mode (production or test).
- `rs`: An instance of RequestSenderI, used for sending custom HTTP requests to the API.
- `mode`: An enumeration indicating whether the client operates in "test" or "prod" mode, affecting the API endpoint.
- `Balance`: A pointer to the Balance structure, providing access to balance-related functionalities.
- `Customers`: A pointer to the Customers structure, enabling operations related to customer management.
- `Prices`: A pointer to the Prices structure, facilitating pricing-related actions.
- `Products`: A pointer to the Products structure, offering methods for product management.
- `PaymentLinks`: A pointer to the PaymentLinks structure, used for generating and managing payment links.
- `Checkouts`: A pointer to the Checkouts structure, providing functionality for handling checkout processes.
- `Webhook`: A pointer to the Webhook structure, allowing the client to manage webhook notifications.

### NewClient Function

```go
func NewClient(apiKey, mode string) (*Client, error)
```

### Parameters

- apiKey: A string representing the API key used for authenticating requests.
- mode: A string that determines the operating mode of the client. Acceptable values are:
  - "prod": Indicates that the client operates in production mode.
  - "test": Indicates that the client operates in test mode.

### Returns

(\*Client, error): A pointer to the initialized Client and an error, if any occurs during initialization.

### Description

The NewClient function initializes a new instance of the Client structure, setting up the necessary configurations based on the provided API key and operational mode.

- The function verifies the mode parameter to determine the appropriate API base URL:
  - If the mode is "prod", it sets the API base URL to the production endpoint.
  - If the mode is "test", it sets the API base URL to the test endpoint.
  - If the mode is invalid, it returns an ErrInvalidMode error.
- It creates a new request sender to facilitate HTTP communications.
- The resulting Client provides access to all functionalities within the Chargily API ecosystem, enabling efficient management of payments and related services.

### Example Usage

```go
client, err := NewClient("your_api_key", "prod")
if err != nil {
    // Handle error
}

// Accessing different functionalities
// client.Balance.Get()
// client.Customers.Crate(), .Update() ...
// client.Prices.Create() ...
// client.Products.Update() ...
// client.PaymentLinks.Create(), ...
// client.Checkouts.GetAll() ...
// client.Webhook.Setup()
```
