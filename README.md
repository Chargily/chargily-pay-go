# [Chargily Pay](https://chargily.com/business/pay "Chargily Pay")™ Gateway - V2.

This Go package provides a client to interact with the Chargily API, allowing you to manage balances, customers, products, prices, checkouts, and payment links.

## About Chargily Pay™ packages

Chargily Pay™ packages/plugins are a collection of open source projects published by Chargily to facilitate the integration of our payment gateway into different programming languages and frameworks. Our goal is to empower developers and businesses by providing easy-to-use tools to seamlessly accept payments.

## API Documentation

For detailed instructions on how to integrate with our API and utilize Chargily Pay™ in your projects, please refer to our [API Documentation](https://dev.chargily.com/pay-v2/introduction).

---

## Navigation

- [Chargily Pay™ Gateway - V2.](#chargily-pay-gateway---v2)
  - [About Chargily Pay™ packages](#about-chargily-pay-packages)
  - [API Documentation](#api-documentation)
  - [Navigation](#navigation)
  - [Installation](#installation)
  - [Usage](#usage)
    - [1- Initialize the Client](#1--initialize-the-client)
  - [2- Retrieve Balance](#2--retrieve-balance)
  - [3- Customers Management](#3--customers-management)
    - [1. Create a New Customer](#1-create-a-new-customer)
    - [2. Update a Customer](#2-update-a-customer)
    - [3. Get a Customer](#3-get-a-customer)
    - [4. Get All Customers](#4-get-all-customers)
  - [4- Products Management](#4--products-management)
    - [1. Create a Product](#1-create-a-product)
    - [2. Update a Product](#2-update-a-product)
    - [3. Get a Product](#3-get-a-product)
    - [4. Get All Products](#4-get-all-products)
    - [5. Delete a Product](#5-delete-a-product)
    - [6. Get Product Prices](#6-get-product-prices)
  - [5- Prices Management](#5--prices-management)
    - [1. Create a Price](#1-create-a-price)
    - [2. Update a Price Metadata](#2-update-a-price-metadata)
    - [4. Get a Price](#4-get-a-price)
    - [5. Get All Prices](#5-get-all-prices)
  - [6- Checkouts Management](#6--checkouts-management)
    - [1. create a checkout](#1-create-a-checkout)
    - [2. Get a Checkout](#2-get-a-checkout)
    - [3. Get All Checkouts](#3-get-all-checkouts)
    - [4. Get Checkout Items](#4-get-checkout-items)
    - [5. Expire a Checkout](#5-expire-a-checkout)
  - [7- Payment Links Management](#7--payment-links-management)
    - [1. Create a Payment Link](#1-create-a-payment-link)
    - [2. Update a Payment Link](#2-update-a-payment-link)
    - [3. Retrieve payment link](#3-retrieve-payment-link)
    - [4. Retrieve All payment links](#4-retrieve-all-payment-links)
    - [5. Retrieve Payment Link Items](#5-retrieve-payment-link-items)
  - [Developers Community](#developers-community)
  - [How to Contribute](#how-to-contribute)
  - [Get in Touch](#get-in-touch)

## Installation

To install the package, use the following command:

```sh
go get github.com/Chargily/chargily-pay-go
```

## Usage

### 1- Initialize the Client

To create a new Chargily API client, you need to provide an API key and the mode (prod for production or test for development):

```go
client, err := chargily.NewClient("your-api-key", "prod")
if err != nil {
    log.Fatalf("Failed to initialize client: %v", err)
}
```

## 2- Retrieve Balance

You can retrieve your account balance using:

```go
balance, err := client.GetBalance()
if err != nil {
    log.Fatalf("Error fetching balance: %v", err)
}
fmt.Printf("Account Balance: %v\n", balance)

```

## 3- Customers Management

### 1. Create a New Customer

```go
// Create a new customer
customerParams := &chargily.CreateCustomerParams{
    Name:  "John Doe",
    Email: "john.doe@example.com",
    Phone: "+1 123 456 7890",
	Address: &chargily.Address{
		Address: "123 Main St",
        City:    "New York",
        Country: "US",
	},
	Metadata: map[string]any{
        "custom_field": "value",
    },
}
customer, err := client.CreateCustomer(customerParams)
if err!= nil {
	fmt.Printf("Error creating customer: %v\n", err)
}
```

### 2. Update a Customer

```go
// update an existing customer
customerParams := &chargily.CreateCustomerParams{
    Name:  "John Doe Doe",
    Email: "john.doe-updated@example.com",
    Phone: "+1 123 456 7891",
    Address: &chargily.Address{
        Address: "123 Main St",
        City:    "New York",
        Country: "US",
    },
	Metadata: map[string]any{
        "custom_field": "value",
    },
}

// Get a customer by ID
customerID := string("your-customer-id") // you should always provide a valid customer ID that already Exists
customer, err := client.UpdateCustomer(customerID,customerParams)
if err!= nil {
    fmt.Printf("Error getting customer: %v\n", err)
    return
}
```

### 3. Get a Customer

```go
customer, err := client.GetCustomer(string("customer_id"))
if err != nil {
    fmt.Println("Error retrieving customer:", err)
    return
}
fmt.Println("Customer:", customer)
```

### 4. Get All Customers

```go
customers, err := client.GetAllCustomers()
if err != nil {
    fmt.Println("Error fetching customers:", err)
    return
}
fmt.Println("Customers:", customers)
```

## 4- Products Management

### 1. Create a Product

```go
// Create a new product
product := &chargily.CreateProductParams{
    Name:           "New Product",
    Description:    "This is a new product",
	Images:         []string{"link1","link2","link3"},
	Metadata:       map[string]any{"key0": "value0","key1":"value1"},
}

newProduct, err := client.CreateProduct(product)
if err!= nil {
    fmt.Printf("Error creating product: %v\n", err)
    return
}
```

### 2. Update a Product

```go
// new data
product := &chargily.CreateProductParams{
    Name:        "updated Product",
    Description: "This is an updates product",
	Images: []string{"link1","link2","link3","link4"},
	Metadata:  map[string]any{"key0": "value0"},
}

// the id of the product to update
id := "your-product-id" // make sure to provide a valid id for the product
// Update the product
updatedProduct, err := client.UpdateProduct(string(id), product)
if err!= nil {
    fmt.Printf("Error updating product: %v\n", err)
    return
}
```

### 3. Get a Product

```go
product, err := client.GetProduct("product_id")
if err != nil {
    fmt.Println("Error retrieving product:", err)
    return
}
// use the product data
fmt.Println("Product:", product)
```

### 4. Get All Products

```go
products, err := client.GetAllProducts()
if err != nil {
    fmt.Println("Error fetching products:", err)
    return
}
fmt.Println("Products:", products)
```

### 5. Delete a Product

```go

//string() method used here to force the id to be always string
err := client.DeleteProduct(string("product_id"))
if err != nil {
    fmt.Println("Error deleting product:", err)
    return
}
fmt.Println("Product deleted successfully.")
```

### 6. Get Product Prices

```go
prices, err := client.GetProductPrices("product_id")
if err != nil {
    fmt.Println("Error retrieving product prices:", err)
    return
}
fmt.Println("Product Prices:", prices)
```

## 5- Prices Management

### 1. Create a Price

```go
// Create a new customer
price := &chargily.ProductPriceParams{
    ProductID: "your-product-id",
    Currency: "USD",
	Amount: 100,
	Metadata: map[string]any{"key":"value"},
}

// create the product price and get back the price object
productPrice, err := client.CreatePrice(price)
if err!= nil {
	fmt.Printf("Error creating price: %v\n", err)
}
```

### 2. Update a Price Metadata

```go
// Create a new customer
price := &chargily.UpdatePriceMetaDataParams{
    Metadata: map[string]any{"key":"val"},
}


//product ID
prodId := "your-product-id"

// update price metadata
productPrice, err := client.UpdatePrice(prodId,price)
if err!= nil {
	fmt.Printf("Error updating price: %v\n", err)
}
```

### 4. Get a Price

```go
price, err := client.GetPrice("product_id")
if err != nil {
    fmt.Println("Error retrieving price:", err)
    return
}
fmt.Println("Price:", price)
```

### 5. Get All Prices

```go
prices, err := client.GetAllPrices()
if err != nil {
    fmt.Println("Error fetching prices:", err)
    return
}
fmt.Println("Prices:", prices)
```

## 6- Checkouts Management

### 1. create a checkout

```go
// Initialize the CheckoutParams struct with adjusted fields.
checkoutParams := &chargily.CheckoutParams{
	Items: items,
	PaymentMethod:   "edahabia",
	SuccessURL:      "https://your-site.com/success",
	FailureURL:      "https://your-site.com/failure",
	WebhookEndpoint: "https://your-site.com/webhook",
	Description:     "Checkout for Order #12345",
	Locale:          "en",
	PercentageDiscount: 10,
}

// Create the checkout
checkoutdata, err := client.CreateCheckout(checkoutParams)
if err!= nil {
    fmt.Printf("Error creating checkout: %v\n", err)
    return
}
```

### 2. Get a Checkout

```go
checkout, err := client.GetCheckout("checkout_id")
if err != nil {
    fmt.Println("Error retrieving checkout:", err)
    return
}
fmt.Println("Checkout:", checkout)
```

### 3. Get All Checkouts

```go
checkouts, err := client.GetAllCheckouts()
if err != nil {
    fmt.Println("Error fetching checkouts:", err)
    return
}
fmt.Println("Checkouts:", checkouts)
```

### 4. Get Checkout Items

```go
items, err := client.GetCheckoutItems("checkout_id")
if err != nil {
    fmt.Println("Error retrieving checkout items:", err)
    return
}
fmt.Println("Checkout Items:", items)
```

### 5. Expire a Checkout

```go
expiredCheckout, err := client.ExpireCheckout("checkout_id")
if err != nil {
    fmt.Println("Error expiring checkout:", err)
    return
}
fmt.Println("Checkout expired:", expiredCheckout)
```

## 7- Payment Links Management

### 1. Create a Payment Link

```go
// Create the payment link parameters
paymentLinkParams := &chargily.CreatePaymentLinkParams{
	Name:                   "Test Order for Payment Link",
	Items:                  itemss,
	AfterCompletionMessage: "Thank you for your order!",
	Locale:                 "en",
	PassFeesToCustomer:     false,
	CollectShippingAddress: 1,
	Metadata: map[string]any{
		"order_id": "order_54321",
		"notes":    "This is a test order for payment link.",
	},
}

// Create the payment link
paymentLink, err := client.CreatePaymentLink(paymentLinkParams)
if err!= nil {
    fmt.Printf("Error creating payment link: %v\n", err)
    return
}
```

### 2. Update a Payment Link

```go
// updated payment link data
paymentLinkParams := &chargily.CreatePaymentLinkParams{
	Name:                   "Test Order for Payment Link",
	Items:                  itemss,
	AfterCompletionMessage: "Thank you for your order!",
	Locale:                 "en",
	PassFeesToCustomer:     true,
	CollectShippingAddress: 0,
	Metadata: map[string]any{
		"order_id": "order_54321",
		"notes":    "This is a test order for payment link.",
	},
}

// the id of payment id to be updated
paymentLinkId:= "your-payment-link-id"

// update
paymentLink, err := client.UpdatePaymentLink(paymentLinkId,paymentLinkParams)
if err!= nil {
    fmt.Printf("Error updating payment link: %v\n", err)
    return
}
```

### 3. Retrieve payment link

```go
// payment link id
paymentLinkId := "your-payment-link-id" // make sure to be a valid payment link id

// retrieve the payment link
paymentLink, err := client.GetPaymentLink(paymentLinkId)
if err!= nil {
    fmt.Printf("Error retrieving payment link by ID: %v\n", err)
    return
}
```

### 4. Retrieve All payment links

```go
// Create a new client instance
client, err := chargily.NewClient(apiKey, mode)

if err!= nil {
    fmt.Printf("Error creating client: %v\n", err)
    return
}

// retrieve all payment links
paymentLinks, err := client.GetAllPaymentLinks()

if err!= nil {
    fmt.Printf("Error retrieving all payment links: %v\n", err)
    return
}
```

### 5. Retrieve Payment Link Items

```go
//payment link ID
paymentLinkId  := "existing-payment-link-id" //must be unique and existing

// retrieve a payment link's items
paymentLinkItems, err := client.GetPaymentLinkItems(string(paymentLinkId))

if err!= nil {
    fmt.Printf("Error getting payment link items: %v\n", err)
    return
}
```

## Developers Community

Join our developer community on Telegram to connect with fellow developers, ask questions, and stay updated on the latest news and developments related to Chargily Pay™ : [Telegram Community](https://chargi.link/PayTelegramCommunity)

## How to Contribute

We welcome contributions of all kinds, whether it's bug fixes, feature enhancements, documentation improvements, or new plugin/package developments. Here's how you can get started:

1. **Fork the Repository:** Click the "Fork" button in the top-right corner of this page to create your own copy of the repository.

2. **Clone the Repository:** Clone your forked repository to your local machine using the following command:

```
git clone https://github.com/Chargily/chargily-pay-go.git
```

3. **Make Changes:** Make your desired changes or additions to the codebase. Be sure to follow our coding standards and guidelines.

4. **Test Your Changes:** Test your changes thoroughly to ensure they work as expected.

5. **Submit a Pull Request:** Once you're satisfied with your changes, submit a pull request back to the main repository. Our team will review your contributions and provide feedback if needed.

## Get in Touch

Have questions or need assistance? Join our developer community on [Telegram](https://chargi.link/PayTelegramCommunity) and connect with fellow developers and our team.

We appreciate your interest in contributing to Chargily Pay™! Together, we can build something amazing.

Happy coding!
