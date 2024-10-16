# Prices Module Documentation

## Overview

The `Prices` struct in the `chargily` package provides methods for managing product prices within the Chargily payment API. This module allows users to create, update, retrieve, and list product prices.

## Prices Structure

### Fields

- **client**: A pointer to the `Client` structure, which holds the API key, endpoint, and other necessary configurations to make API requests.

## Methods

### Create

#### Parameters

- **productPrice**: A pointer to a `models.ProductPriceParams` structure containing the details of the product price to be created.

#### Returns

- **(\*models.ProductPrice, error)**: A pointer to the newly created `models.ProductPrice` object and an error, if any occurs during the request.

#### Description

The `Create` method sends a request to the Chargily API to create a price for a specific product with the provided details. It constructs the API request and parses the response into a `models.ProductPrice` object.

#### Example

```go
func CreatePrice(client *chargily.Client, productID string) {
	// Define the parameters for the new price.
	bodyRequestPrice := &models.ProductPriceParams{
		Amount:    10000,
		Currency:  "dzd",
		ProductID: productID, // Product ID to associate with this price
		Metadata:  map[string]any{"key": "value"},
	}

	// Call the Create method on the Prices service to create the price.
	price, err := client.Prices.Create(bodyRequestPrice)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating price: %v\n", err)
		return
	}
	// Output the details of the successfully created price.
	fmt.Printf("Price created successfully: %+v\n", price)
}
```

---

### Update

#### Parameters

- **productId**: A string representing the unique identifier of the product whose price metadata is to be updated.
- **Data**: A pointer to a `models.UpdatePriceMetaDataParams` structure containing the updated price metadata.

#### Returns

- **(\*models.ProductPrice, error)**: A pointer to the updated `models.ProductPrice` object and an error, if any occurs during the request.

#### Description

The `Update` method sends a request to the Chargily API to update the metadata of a product price identified by `productId`. It constructs the API request with the new metadata and returns the updated product price object.

#### Example

```go
func UpdatePrice(client *chargily.Client, priceID string) {
	// Define the parameters for updating the price metadata.
	data := &models.UpdatePriceMetaDataParams{
		Metadata: map[string]any{"new_key": "new_value", "key3": "keykey"},
	}

	// Call the Update method on the Prices service to update the price.
	price, err := client.Prices.Update(priceID, data)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating price: %v\n", err)
		return
	}
	// Output the details of the successfully updated price.
	fmt.Printf("Price updated successfully: %+v\n", price)
}
```

---

### Get

#### Parameters

- **productId**: A string representing the unique identifier of the product whose price is to be retrieved.

#### Returns

- **(\*models.ProductPrice, error)**: A pointer to the retrieved `models.ProductPrice` object and an error, if any occurs during the request.

#### Description

The `Get` method retrieves the price details of a specific product identified by `productId`. It sends a GET request to the API and parses the response into a `models.ProductPrice` object.

#### Example

```go
func RetrievePrice(client *chargily.Client, priceID string) {
	// Retrieve the price by its ID.
	price, err := client.Prices.Get(priceID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving price: %v\n", err)
		return
	}
	// Output the details of the retrieved price.
	fmt.Printf("Price retrieved successfully: %+v\n", price)
}

```

---

### GetAll

#### Returns

- **(\*models.RetrieveAll[models.ProductPrice], error)**: A pointer to a `models.RetrieveAll[models.ProductPrice]` structure containing an array of all product prices and an error, if any occurs during the request.

#### Description

The `GetAll` method retrieves a list of all available product prices stored in the Chargily system. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.ProductPrice]` structure, which includes an array of product price objects.

#### Example

```go
func GetAllPrices(client *chargily.Client) {
	// Call the GetAll method on the Prices service to retrieve all prices.
	prices, err := client.Prices.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all prices: %v\n", err)
		return
	}
	// Output the list of successfully retrieved prices.
	fmt.Printf("All prices retrieved successfully: %+v\n", prices)
}
```
