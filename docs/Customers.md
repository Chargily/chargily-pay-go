# Customers Module Documentation

## Overview

The `Customers` struct in the `chargily` package provides methods for managing customer-related functionalities within the Chargily payment API. This module allows the creation, updating, retrieval, and deletion of customer records, enabling seamless customer management.

## Customers Structure

### Definition

```go
type Customers struct {
    client *Client // Reference to the Client for API interactions
}
```

### Fields

client: A pointer to the Client structure, which holds the API key, endpoint, and other necessary configurations to make API requests.

## Methods

### Create

```go
func (c *Customers) Create(customer *models.CreateCustomerParams) (*models.Customer, error)
```

#### Parameters

`customer`: A pointer to a `models.CreateCustomerParams `structure containing the details of the customer to be created.

#### Returns

(\*models.Customer, error): A pointer to the newly created models.Customer object and an error, if any occurs during the request.

#### Description

The Create method sends a request to the Chargily API to create a new customer with the provided details. It constructs the API request and parses the response into a models.Customer object.

#### Example

```go
func CreateCustomer(client * chargily.Client) {
	// Param example for a clear example using types.
	customerParams := &models.CreateCustomerParams{
		Name:  "John Doe", // Customer's name
		Email: "john.doe@example.com", // Customer's email
		Phone: "+1234567890", // Customer's phone number
		Address: &models.Address{
			Address: "123 Main St", // Street address
			State:   "NY",         // State
			Country: "US",         // Country
		},
	}

	// Call the Create method on the Customers service to create the customer.
	customer, err := client.Customers.Create(customerParams)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating customer: %v\n", err)
		return
	}
	// Output the successfully created customer details.
	fmt.Printf("Customer created successfully: %+v\n", *customer)
}
```

### Update

```go
func (c *Customers) Update(customerID string, customer *models.CreateCustomerParams) (\*models.Customer, error)
```

#### Parameters

`customerID`: A string representing the unique identifier of the customer to be updated.
`customer`: A pointer to a `models.CreateCustomerParams` structure containing the updated customer details.

#### Returns

(\*models.Customer, error): A pointer to the updated models.Customer object and an error, if any occurs during the request.

#### Description

The Update method sends a request to the Chargily API to update the information of an existing customer specified by customerID. It constructs the API request with the new customer data and returns the updated customer object.

#### Example

```go
func UpdateCustomer(customerID string, client * chargily.Client) {

	// Define the parameters for updating the customer.
	customerParams := &models.CreateCustomerParams{
		Email: "john.updated.doe@example.com", // Updated email
		Address: &models.Address{
			Address: "1234 Main St", // Updated street address
			State:   "NYC",         // State
			Country: "USA",         // Country
		},
	}

	// Call the Update method on the Customers service to update the customer.
	customer, err := client.Customers.Update(customerID, customerParams)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating customer: %v\n", err)
		return
	}
	// Output the successfully updated customer details.
	fmt.Printf("Customer updated successfully: %+v\n", *customer)
}
```

### Get

```go
func (c *Customers) Get(customerID string) (*models.Customer, error)
```

#### Parameters

`customerID`: A string representing the unique identifier of the customer to be retrieved.

#### Returns

(\*models.Customer, error): A pointer to the retrieved `models.Customer` object and an error, if any occurs during the request.

#### Description

The Get method retrieves the details of a specific customer identified by customerID. It sends a GET request to the API and parses the response into a models.Customer object.

#### Example

```go
func GetCustomer(customerID string, client *chargily.Client) {

	// Call the Get method on the Customers service to retrieve the customer by ID.
	customer, err := client.Customers.Get(customerID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error getting customer: %v\n", err)
		return
	}
	// Output the successfully retrieved customer details.
	fmt.Printf("Customer retrieved successfully: %+v\n", *customer)
}
```

### Delete

```go
func (c \*Customers) Delete(customerID string) error
```

#### Parameters

`customerID`: A string representing the unique identifier of the customer to be deleted.

#### Returns

error: An error, if any occurs during the request.

#### Description

The Delete method sends a request to the Chargily API to delete a specific customer identified by customerID. If the request is successful, it returns nil; otherwise, it returns an error.

#### Example

```go
func DeleteCustomer(customerID string , client * chargily.Client) {
	// Call the Delete method on the Customers service to delete the customer by ID.
	if err := client.Customers.Delete(customerID); err != nil {
		// Handle deletion error.
		fmt.Printf("Error deleting customer: %v\n", err)
		return
	}
	// Output a success message for deletion.
	fmt.Println("Customer deleted successfully.")
}
```

### GetAll

```go
func (c *Customers) GetAll() (*models.RetrieveAll[models.Customer], error)
```

#### Returns

(\*models.RetrieveAll[models.Customer], error): A pointer to a models.RetrieveAll[models.Customer] structure containing an array of all customers and an error, if any occurs during the request.

#### Description

The GetAll method retrieves a list of all customers stored in the Chargily system. It sends a GET request to the API and parses the response into a models.RetrieveAll[models.Customer] structure, which includes an array of customer objects.

#### Example

```go
func GetAllCustomers(client * chargily.Client) {
	// Call the GetAll method on the Customers service to retrieve all customers.
	customers, err := client.Customers.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error getting all customers: %v\n", err)
		return
	}
	// Output the successfully retrieved list of customers.
	fmt.Printf("All customers retrieved successfully: %+v\n", customers)
}
```
