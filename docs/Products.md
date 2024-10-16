# Products Module Documentation

## Overview

The `Products` struct in the `chargily` package provides methods for managing product-related functionalities within the Chargily payment API. This module allows users to create, update, retrieve, delete products, and manage product prices.

## Products Structure

### Fields

- **client**: A pointer to the `Client` structure, which holds the API key, endpoint, and other necessary configurations to make API requests.

## Methods

### Create

#### Parameters

- **product**: A pointer to a `models.CreateProductParams` structure containing the details of the product to be created.

#### Returns

- **(\*models.Product, error)**: A pointer to the newly created `models.Product` object and an error, if any occurs during the request.

#### Description

The `Create` method sends a request to the Chargily API to create a new product with the provided details. It constructs the API request and parses the response into a `models.Product` object.

#### Example

```go
func CreateProduct(client *chargily.Client) {
	// Define the parameters for the new product.
	bodyRequestProduct := &models.CreateProductParams{
		Name:        "Test Product",
		Description: "This is a test product",
		Images:      []string{"valid-image-link"},
		Metadata:    map[string]any{"key": "value"},
	}

	// Call the Create method on the Products service to create the product.
	product, err := client.Products.Create(bodyRequestProduct)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating product: %v\n", err)
		return
	}
	// Output the details of the successfully created product.
	fmt.Printf("Product created successfully: %+v\n", product)
}
```

---

### Update

#### Parameters

- **productId**: A string representing the unique identifier of the product to be updated.
- **product**: A pointer to a `models.CreateProductParams` structure containing the updated product details.

#### Returns

- **(\*models.Product, error)**: A pointer to the updated `models.Product` object and an error, if any occurs during the request.

#### Description

The `Update` method sends a request to the Chargily API to update the information of an existing product specified by `productId`. It constructs the API request with the new product data and returns the updated product object.

#### Example

```go
func UpdateProduct(client *chargily.Client, productID string) {
	// Define the parameters for updating the product.
	bodyRequestProduct := &models.CreateProductParams{
		Name:        "Test Product of the Update",
		Description: "This is an updated test product",
		Images:      []string{"valid-image-link"},
		Metadata:    map[string]any{"key": "value"},
	}

	// Call the Update method on the Products service to update the product.
	product, err := client.Products.Update(productID, bodyRequestProduct)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating product: %v\n", err)
		return
	}
	// Output the details of the successfully updated product.
	fmt.Printf("Product updated successfully: %+v\n", product)
}
```

---

### Get

#### Parameters

- **productId**: A string representing the unique identifier of the product to be retrieved.

#### Returns

- **(\*models.Product, error)**: A pointer to the retrieved `models.Product` object and an error, if any occurs during the request.

#### Description

The `Get` method retrieves the details of a specific product identified by `productId`. It sends a GET request to the API and parses the response into a `models.Product` object.

---

### GetAll

#### Returns

- **(\*models.RetrieveAll[models.Product], error)**: A pointer to a `models.RetrieveAll[models.Product]` structure containing an array of all products and an error, if any occurs during the request.

#### Description

The `GetAll` method retrieves a list of all products stored in the Chargily system. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.Product]` structure, which includes an array of product objects.

#### Example

```go
func GetAllProducts(client *chargily.Client) {
	// Call the GetAll method on the Products service to retrieve all products.
	products, err := client.Products.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all products: %v\n", err)
		return
	}
	// Output the list of successfully retrieved products.
	fmt.Printf("All products retrieved successfully: %+v\n", products)
}
```

---

### Delete

#### Parameters

- **productId**: A string representing the unique identifier of the product to be deleted.

#### Returns

- **error**: An error, if any occurs during the request.

#### Description

The `Delete` method sends a request to the Chargily API to delete a specific product identified by `productId`. If the request is successful, it returns `nil`; otherwise, it returns an error.

#### Example

```go
func DeleteProduct(client *chargily.Client, productID string) {
	// Call the Delete method on the Products service to delete the product by ID.
	if err := client.Products.Delete(productID); err != nil {
		// Handle deletion error.
		fmt.Printf("Error deleting product: %v\n", err)
		return
	}
	// Output a success message for deletion.
	fmt.Println("Product deleted successfully.")
}
```

---

### GetPrices

#### Parameters

- **productId**: A string representing the unique identifier of the product whose prices are to be retrieved.

#### Returns

- **(\*models.RetrieveAll[models.ProductPrice], error)**: A pointer to a `models.RetrieveAll[models.ProductPrice]` structure containing the prices of the specified product and an error, if any occurs during the request.

#### Description

The `GetPrices` method retrieves the prices associated with a specific product identified by `productId`. It sends a GET request to the API and parses the response into a `models.RetrieveAll[models.ProductPrice]` structure, which includes an array of product prices.

#### Example

```go
func GetProductPrices(productId string, client * chargily.Client){
    // Get a product price
    prices, err := client.Products.GetPrices(productId)
    if err!= nil {
        fmt.Println("Error retrieving product price:", err)
    }

    fmt.Printf("Product prices : %v", prices)
}
```
