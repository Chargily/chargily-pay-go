# Checkouts Module Documentation

## Overview

The `Checkouts` struct in the `chargily` package provides methods for managing checkout functionalities within the Chargily payment API. This module allows users to create, retrieve, and manage checkouts and their associated items.

## Checkouts Structure

### Fields

- **client**: A pointer to the `Client` structure, which holds the API key, endpoint, and other necessary configurations to make API requests.

## Methods

### Create

#### Parameters

- **checkout**: A pointer to a `models.CheckoutParams` structure containing the details needed to create a new checkout.

#### Returns

- **(\*models.Checkout, error)**: A pointer to the newly created `models.Checkout` object and an error, if any occurs during the request.

#### Description

The `Create` method sends a request to the Chargily API to create a new checkout with the provided details. It constructs the API request and parses the response into a `models.Checkout` object.

#### Example

```go
func CreateCheckout(client *chargily.Client , id string) {
	// Create example items to be added to the checkout.
	items := []models.CItems{
		{
			Price:    id,
			Quantity: 2,
		},
	}

	// Initialize the CheckoutParams struct with required fields.
	checkout := &models.CheckoutParams{
		Items:            items,
		PaymentMethod:    "edahabia",
		SuccessURL:       "https://your-site.com/success",
		FailureURL:       "https://your-site.com/failure",
		WebhookEndpoint:  "https://your-site.com/webhook",
		Description:      "Checkout for Order #12345",
		Locale:           "en",
		PercentageDiscount: 10,
	}

	// Create the checkout using the client.
	checkoutData, err := client.Checkouts.Create(checkout)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating checkout: %v\n", err)
		return
	}

	// Output the details of the created checkout.
	fmt.Printf("Checkout created successfully: %+v\n", checkoutData)
	fmt.Printf("Checkout Parameters: %+v\n", checkout)
}
```

---

### Get

#### Parameters

- **checkoutId**: A string representing the unique identifier of the checkout to be retrieved.

#### Returns

- **(\*models.Checkout, error)**: A pointer to the retrieved `models.Checkout` object and an error, if any occurs during the request.

#### Description

The `Get` method retrieves the details of a specific checkout identified by `checkoutId`. It sends a GET request to the API and parses the response into a `models.Checkout` object.

#### Example

```go
func RetrieveCheckout(client *chargily.Client, checkoutID string) {
	// Retrieve the checkout using its ID.
	checkoutData, err := client.Checkouts.Get(checkoutID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving checkout: %v\n", err)
		return
	}
	// Output the details of the retrieved checkout.
	fmt.Printf("Checkout retrieved successfully: %+v\n", checkoutData)
}
```

---

### GetAll

#### Returns

- **(\*models.RetrieveAll[models.Checkout], error)**: A pointer to a `models.RetrieveAll[models.Checkout]` structure containing an array of all checkouts and an error, if any occurs during the request.

#### Description

The `GetAll` method retrieves a list of all available checkouts stored in the Chargily system. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.Checkout]` structure, which includes an array of checkout objects.

#### Example

```go
func ListAllCheckouts(client *chargily.Client) {
	// Call the GetAll method on the Checkouts service to retrieve all checkouts.
	checkouts, err := client.Checkouts.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all checkouts: %v\n", err)
		return
	}
	// Output the list of successfully retrieved checkouts.
	fmt.Printf("All checkouts retrieved successfully: %+v\n", checkouts)
}
```

---

### GetItems

#### Parameters

- **checkoutId**: A string representing the unique identifier of the checkout whose items are to be retrieved.

#### Returns

- **(\*models.RetrieveAll[models.CheckoutItems], error)**: A pointer to a `models.RetrieveAll[models.CheckoutItems]` structure containing an array of items associated with the checkout and an error, if any occurs during the request.

#### Description

The `GetItems` method retrieves the items associated with a specific checkout identified by `checkoutId`. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.CheckoutItems]` structure, which includes an array of item data.

#### Example

```go
func RetrieveCheckoutItems(client *chargily.Client, checkoutID string) {
	// Retrieve the items from the checkout using its ID.
	checkoutItems, err := client.Checkouts.GetItems(checkoutID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving checkout items: %v\n", err)
		return
	}
	// Output the details of the retrieved checkout items.
	fmt.Printf("Checkout items retrieved successfully: %+v\n", checkoutItems)
}
```

---

### Expire

#### Parameters

- **checkoutId**: A string representing the unique identifier of the checkout to be expired.

#### Returns

- **(\*models.Checkout, error)**: A pointer to the expired `models.Checkout` object and an error, if any occurs during the request.

#### Description

The `Expire` method sends a request to the Chargily API to expire a checkout identified by `checkoutId`. It constructs the API request and returns the corresponding checkout object if the expiration is successful.

#### Example

```go
func CancelCheckout(client *chargily.Client, checkoutID string) {
	// Call the Expire method on the Checkouts service to cancel the checkout.
	checkoutEX, err := client.Checkouts.Expire(checkoutID)
	if err != nil {
		// Handle cancellation error.
		fmt.Printf("Error canceling checkout: %v\n", err)
		return
	}
	// Output the details of the successfully canceled checkout.
	fmt.Printf("Checkout canceled successfully: %+v\n", checkoutEX)
}
```
