# Payment Links Module Documentation

## Overview

The `PaymentLinks` struct in the `chargily` package provides methods for managing payment links within the Chargily payment API. This module allows users to create, update, retrieve, and list payment links.

## PaymentLinks Structure

### Fields

- **client**: A pointer to the `Client` structure, which holds the API key, endpoint, and other necessary configurations to make API requests.

## Methods

### Create

#### Parameters

- **paymentLink**: A pointer to a `models.CreatePaymentLinkParams` structure containing the details of the payment link to be created.

#### Returns

- **(\*models.PaymentLink, error)**: A pointer to the newly created `models.PaymentLink` object and an error, if any occurs during the request.

#### Description

The `Create` method sends a request to the Chargily API to create a payment link with the provided details. It constructs the API request and parses the response into a `models.PaymentLink` object.

#### Example

```go
func CreatePaymentLink(client *chargily.Client, id string) {
	// Create items to be added to the payment link
	items := []models.PItems{
		{
			Price:              id,
			Quantity:           2,
			AdjustableQuantity: true,
		},
	}

	// Create the payment link parameters
	paymentLinkParams := &models.CreatePaymentLinkParams{
		Name:                   "Test Order for Payment Link",
		Items:                  items,
		AfterCompletionMessage: "message",
		Locale:                 "en",
		PassFeesToCustomer:     false,
		CollectShippingAddress: 1,
		Metadata: map[string]any{
			"order_id": "order_54321",
			"notes":    "This is a test order for payment link.",
		},
	}

	// Create the payment link using the client.
	paymentLink, err := client.PaymentLinks.Create(paymentLinkParams)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating payment link: %v\n", err)
		return
	}
	// Output the details of the created payment link.
	fmt.Printf("Payment link created successfully: %+v\n", paymentLink)
}
```

---

### Update

#### Parameters

- **paymentLinkId**: A string representing the unique identifier of the payment link to be updated.
- **paymentLink**: A pointer to a `models.CreatePaymentLinkParams` structure containing the updated payment link details.

#### Returns

- **(\*models.PaymentLink, error)**: A pointer to the updated `models.PaymentLink` object and an error, if any occurs during the request.

#### Description

The `Update` method sends a request to the Chargily API to update an existing payment link identified by `paymentLinkId`. It constructs the API request with the new payment link data and returns the updated payment link object.

#### Example

```go
func UpdatePaymentLink(client *chargily.Client, paymentLinkID string) {
	// Prepare updated payment link parameters.
	updatedPaymentLinkParams := &models.CreatePaymentLinkParams{
		Name:                   "Updated Test Order for Payment Link",
		Items:                  []models.PItems{
			{
				Price:              paymentLinkID,
				Quantity:           2,
				AdjustableQuantity: true,
			},
		},
		AfterCompletionMessage: "Thank you for your updated order!",
		Locale:                 "en",
		PassFeesToCustomer:     false,
		CollectShippingAddress: 0,
		Metadata: map[string]any{
			"order_id": "updated_order_54321",
			"notes":    "This is an updated test order for payment link.",
		},
	}

	// Update the payment link using the client.
	updatedPaymentLink, err := client.PaymentLinks.Update(paymentLinkID, updatedPaymentLinkParams)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating payment link: %v\n", err)
		return
	}
	// Output the details of the updated payment link.
	fmt.Printf("Payment link updated successfully: %+v\n", updatedPaymentLink)
}
```

---

### Get

#### Parameters

- **paymentLinkId**: A string representing the unique identifier of the payment link to be retrieved.

#### Returns

- **(\*models.PaymentLink, error)**: A pointer to the retrieved `models.PaymentLink` object and an error, if any occurs during the request.

#### Description

The `Get` method retrieves the details of a specific payment link identified by `paymentLinkId`. It sends a GET request to the API and parses the response into a `models.PaymentLink` object.

#### Example

```go
func RetrievePaymentLink(client *chargily.Client, paymentLinkID string) {
	// Retrieve the payment link using its ID.
	paymentLinkData, err := client.PaymentLinks.Get(paymentLinkID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving payment link: %v\n", err)
		return
	}
	// Output the details of the retrieved payment link.
	fmt.Printf("Payment link retrieved successfully: %+v\n", paymentLinkData)
}
```

---

### GetAll

#### Returns

- **(\*models.RetrieveAll[models.PaymentLink], error)**: A pointer to a `models.RetrieveAll[models.PaymentLink]` structure containing an array of all payment links and an error, if any occurs during the request.

#### Description

The `GetAll` method retrieves a list of all available payment links stored in the Chargily system. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.PaymentLink]` structure, which includes an array of payment link objects.

#### Example

```go
func ListAllPaymentLinks(client *chargily.Client) {
	// Call the GetAll method on the PaymentLinks service to retrieve all payment links.
	paymentLinks, err := client.PaymentLinks.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all payment links: %v\n", err)
		return
	}
	// Output the list of successfully retrieved payment links.
	fmt.Printf("All payment links retrieved successfully: %+v\n", paymentLinks)
}
```

---

### GetItems

#### Parameters

- **productId**: A string representing the unique identifier of the payment link whose items are to be retrieved.

#### Returns

- **(\*models.RetrieveAll[models.PItemsData], error)**: A pointer to a `models.RetrieveAll[models.PItemsData]` structure containing an array of items associated with the payment link and an error, if any occurs during the request.

#### Description

The `GetItems` method retrieves the items associated with a specific payment link identified by `productId`. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.PItemsData]` structure, which includes an array of item data.

#### Example

```go
func RetrievePaymentLinkItems(client *chargily.Client, paymentLinkID string) {
	// Retrieve the items from the payment link using its ID.
	paymentLinkItems, err := client.PaymentLinks.GetItems(paymentLinkID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving payment link items: %v\n", err)
		return
	}
	// Output the details of the retrieved payment link items.
	fmt.Printf("Payment link items retrieved successfully: %+v\n", paymentLinkItems)
}
```
