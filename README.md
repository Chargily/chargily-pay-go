# [Chargily Pay](https://chargily.com/business/pay "Chargily Pay")™ Gateway - V2.

This Go package provides a client to interact with the Chargily API, allowing you to manage balances, customers, products, prices, checkouts, and payment links.

### About Chargily Pay™ packages

Chargily Pay™ packages/plugins are a collection of open source projects published by Chargily to facilitate the integration of our payment gateway into different programming languages and frameworks. Our goal is to empower developers and businesses by providing easy-to-use tools to seamlessly accept payments.

### API Documentation

For detailed instructions on how to integrate with our API and utilize Chargily Pay™ in your projects, please refer to our [API Documentation](https://dev.chargily.com/pay-v2/introduction).

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Usage](#usage)
4. [Documentation](#documentation)

## Overview

This Golang SDK provides a seamless integration with the [Chargily API](https://chargily.com) for interacting with their payment gateway services. It simplifies the process of integrating Chargily's e-payment solutions into your Go applications, allowing you to handle payments and transactions with ease.

### Use Cases:

- E-commerce websites that need to process payments through Chargily’s gateway.
- SaaS platforms requiring integration with a local payment processor.
- Any Go-based application that needs to handle payment processing in the Algerian market.

The SDK abstracts much of the complexity of interacting with the Chargily API, providing a simple and easy-to-use interface for developers to build payment functionalities into their Go applications.

## Features

Below are the key functionalities available through the client:

### 1. Payment Checkout

- **Create Checkout**: Create a new checkout session for processing payments.
- **Cancel Checkout**: Cancel an ongoing checkout.
- **Get Checkout**: Retrieve details of a specific checkout.
- **List All Checkouts**: Get a list of all checkout sessions.
- **Get Checkout Items**: Fetch items associated with a specific checkout.

### 2. Customer Management

- **Create Customer**: Add a new customer to the system.
- **Get Customer**: Retrieve information about a specific customer.
- **List All Customers**: Fetch all customers from the system.
- **Update Customer**: Modify existing customer details.
- **Delete Customer**: Remove a customer from the system.

### 3. Product Management

- **Create Product**: Add a new product to the inventory.
- **Get Product**: Retrieve details of a specific product.
- **Retrieve Product Prices**: Retrieves all the prices of a product.
- **List All Products**: Fetch a list of all products in the system.
- **Update Product**: Update existing product details.
- **Delete Product**: Remove a product from the inventory.

### 4. Payment Links

- **Create Payment Link**: Generate payment links that can be shared and directs your customers to a hosted payment page.
- **Get Payment Link**: Get a specific payment link .
- **Get All Payment Links**: Get all payment links.
- **Get Payment Link Items**: Get all payment link associated items.
- **Update Payment Link**: Updates the information of a payment link.

### 5. Price Management

- **Get Price**: Retrieve pricing details for a specific product.
- **List All Prices**: Fetch all prices listed in the system.
- **Update Price**: Updates the information of a price of an existing product.

### 6. Webhook Management

- **Setup Webhook**: Configure webhooks to receive updates from Chargily's server.
- **Verify Webhook Signature**: Ensure security by verifying Chargily webhook signatures.

## Usage

To start using the Chargily SDK in your Go project, you can install it via `go get`:

```sh
go get github.com/Chargily/chargily-pay-go
```

### Initialization

First, import the SDK into your project:

```go
import (
    //to import the client and the functionalities
    "github.com/Chargily/chargily-pay-go/pkg/chargily"

    //to import the models (types) and use them in the project (params, types, generic response).
    "github.com/Chargily/chargily-pay-go/pkg/models"
)
```

## Documentation

For more detailed information about specific features and usage, check out the documentation:

- [API Client Setup](./docs/Client.md): Detailed guide on initializing the client and interacting with the API.
- [Customer Management](./docs/Customers.md): Learn how to create, update, delete, and manage customers.
- [Checkouts Management](./docs/Checkouts.md): Find out how to create, cancel, and retrieve checkouts.
- [Payment Links Management](./docs/PaymentLinks.md): Find out how to fully manage payment links.
- [Product Management](./docs/Products.md): Documentation on how to create, update, and manage products and their prices.
- [Prices Management](./docs/Prices.md): Learn how to set products prices,update or retrieve.
- [Webhook Integration](./docs/Webhook.md): Learn how to set up and verify webhooks to receive real-time notifications.

### Additional Resources:

- [SDK Examples](./pkg/examples): Practical examples on using the SDK in real-world scenarios.

For any additional questions or further assistance, feel free to reach out via the project's [issues page](#) and Never skip the Api Docs at [Chargily API](https://dev.chargily.com/pay-v2/api-reference/).

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

```

```
