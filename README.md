# Welcome to Go Package Repository
# for [Chargily Pay](https://chargily.com/business/pay "Chargily Pay")™ Gateway - V2.

Thank you for your interest in Go Package of Chargily Pay™, an open source project by Chargily, a leading fintech company in Algeria specializing in payment solutions and  e-commerce facilitating, this Package is providing the easiest and free way to integrate e-payment API through widespread payment methods in Algeria such as EDAHABIA (Algerie Post) and CIB (SATIM) into your digital projects.

This package is developed by **Anas Bahaa Eddine ([anes011](https://github.com/anes011))** and is open to contributions from developers like you.

## Installation
```bash
go get github.com/Chargily/chargily-pay-go
```

## Usage
```go
package main

import "github.com/Chargily/chargily-pay-go"

func main() {
    // Initialize chargily client
    chargilypaygo.NewChargily("<MODE>", "<API-KEY>")

    // Example
    client, err := chargilypaygo.NewChargily("test", "test_sk_123456789abcdefg")
    if err != nil {
        log.Fatalf("Error initializing chargily client, err: %s", err.Error())
    }

    // Get Balance
    resp, err := client.GetBalance()
    if err != nil {
        log.Fatalf("Could not get balance, err: %s", err.Error())
    }

    fmt.Println(string(resp))

    // Create Customer
    customerRequest := map[string]any{
        "name": "Customer Example",
        "email": "customer@example.com",
        "phone": "+213 123 456 789",
        "address": map[string]any{
            "country": "Algeria",
            "state": "Algiers"
        },
    }

    resp, err := client.CreateCustomer(customerRequest)
    if err != nil {
        log.Fatalf("Could not create customer, err: %s", err.Error())
    }

    fmt.Println(string(resp))

    // Update Customer
    customerRequest := map[string]any{
        "name": "Updated Customer Example",
        "email": "customer@example.com",
        "phone": "+213 123 456 789",
        "address": map[string]any{
            "country": "Algeria",
            "state": "Algiers"
        },
    }

    resp, err := client.UpdateCustomer("<CUSTOMER_ID>", customerRequest)
    if err != nil {
        log.Fatalf("Could not update customer, err: %s", err.Error())
    }

    fmt.Println(string(resp))
    
    // Retrieve a Customer
    resp, err := client.RetrieveCustomer("<CUSTOMER_ID>")
    if err != nil {
        log.Fatalf("Could not retrieve customer, err: %s", err.Error())
    }

    fmt.Println(string(resp))
    
    // List all Customers
    resp, err := client.ListCustomers("<QUANTITY>")
    if err != nil {
        log.Fatalf("Could not list customers, err: %s", err.Error())
    }

    fmt.Println(string(resp))
    
    // Delete a Customer
    resp, err := client.DeleteCustomer("<CUSTOMER_ID>")
    if err != nil {
        log.Fatalf("Could not delete customer, err: %s", err.Error())
    }

    fmt.Println(string(resp))

    // Create a Product
    productRequest := map[string]any{
        "name": "Product Example",
        "description": "This is a product",
        "images": []string{"https://image.png", "https://image2.png"},
    }

    resp, err := client.CreateProduct(productRequest)
    if err != nil {
        log.Fatalf("Could not create product, err: %s", err.Error())
    }

    fmt.Println(string(resp))

    // Create a Price
    priceRequest := map[string]any{
        "amount": 2000.56,
        "currency": "dzd",
        "product_id": "<PRODUCT_ID>",
    }

    resp, err := client.CreatePrice(priceRequest)
    if err != nil {
        log.Fatalf("Could not create price, err: %s", err.Error())
    }

    fmt.Println(string(resp))
    
    // Create a Checkout
    checkoutRequest := map[string]any{
        "amount": 2000.56,
        "currency": "dzd",
        "success_url": "https://success.com"
    }

    resp, err := client.CreateCheckout(checkoutRequest)
    if err != nil {
        log.Fatalf("Could not create checkout, err: %s", err.Error())
    }

    fmt.Println(string(resp))
    
    // Create a Payment Link
    paymentLinkRequest := map[string]any{
        "name": "any name",
        "items": []map[string]any{
            {
                "price": "<PRICE_ID>",
                "quantity": 1,
            },
        },
    }

    resp, err := client.CreatePaymentLink(paymentLinkRequest)
    if err != nil {
        log.Fatalf("Could not create payment link, err: %s", err.Error())
    }

    fmt.Println(string(resp))

    // Access certain fields in the response
    checkoutRequest := map[string]any{
        "amount": 2000.56,
        "currency": "dzd",
        "success_url": "https://success.com"
    }

    resp, err := client.CreateCheckout(checkoutRequest)
    if err != nil {
        log.Fatalf("Could not create checkout, err: %s", err.Error())
    }

    data := map[string]any{}

    err = json.Unmarshal(resp, &data)
    if err != nil {
        log.Fatalf("Error unmarshaling response: %s", err.Error())
    }

    fmt.Println(data["checkout_url"])
}
```

## About Chargily Pay™ packages

Chargily Pay™ packages/plugins are a collection of open source projects published by Chargily to facilitate the integration of our payment gateway into different programming languages and frameworks. Our goal is to empower developers and businesses by providing easy-to-use tools to seamlessly accept payments.

## API Documentation

For detailed instructions on how to integrate with our API and utilize Chargily Pay™ in your projects, please refer to our [API Documentation](https://dev.chargily.com/pay-v2/introduction). 

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

