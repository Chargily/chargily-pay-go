# Balance Struct Documentation

The `Balance` struct in the `chargily` package is used to retrieve the balance from the Chargily API.

## Method: `Get`

This method retrieves the current balance from the Chargily API.

### Syntax

```go
func (b *Balance) Get() (*models.Balance, error)
```

### Description

The Get method sends a GET request to the Chargily API to retrieve the current balance.
It returns the balance data in the form of a models.Balance object, along with any error that occurred during the request.

### Parameters

No params passed

### Returns

\*models.Balance: The current balance as a Balance object.
error: If the request fails, an error is returned.

### Example Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/Chargily/chargily-pay-go/pkg/chargily"
)

func main() {
    // Initialize the API client
    client := chargily.NewClient("API_KEY")


    // Retrieve the balance
    balance, err := client.Balance.Get()
    if err != nil {
        log.Fatalf("Error retrieving balance: %v", err)
    }

    // Print the balance
    fmt.Printf("Current Balance: %+v\n", balance)
}
```

### Error Handling

If an error occurs during the request, the Get method will return an error object. You should always check the error and handle it appropriately.
