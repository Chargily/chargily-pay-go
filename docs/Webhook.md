# Webhook Module Documentation

## Overview

The `Webhook` struct in the `chargily` package is designed to handle webhook events from the Chargily payment API. It provides methods for setting up handlers to process incoming webhook notifications securely, ensuring that only valid requests are processed.

## Webhook Structure

### Fields

- **client**: A pointer to the `Client` structure, which contains the API key and endpoint necessary for making API requests and verifying webhook signatures.

## Methods

### SetupHandler

#### Parameters

- **path**: A string representing the HTTP path where the webhook notifications will be received.
- **handler**: A function that takes an `eventType` (string) and a `models.WebhookEvent`. This function will be called when a webhook event is received.

#### Returns

- **void**

#### Description

The `SetupHandler` method sets up an HTTP handler for incoming webhook notifications. It performs the following tasks:

1. Validates the presence of a signature in the request header.
2. Reads and verifies the payload using the provided signature.
3. Parses the JSON payload into a `models.WebhookEvent` object.
4. Identifies the event type and invokes the user-defined event handler with the event type and parsed event data.
5. Sends a 200 OK response back to the Chargily server upon successful processing of the webhook.

#### Example

```go
func SetupWebhook(client * chargily.Client){
	// Setup the webhook endpoint to listen for webhook events on `/webhook`
	client.Webhook.SetupHandler("/webhook", handleEvent)
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// Add your custom event handler function that processes events you can see examples folder
func handleEvent(eventType string, event models.WebhookEvent) {...}
```

---

### VerifySignature

#### Parameters

- **payload**: A byte slice containing the raw data received in the webhook request.
- **signature**: A string containing the signature from the request header.

#### Returns

- **error**: An error object if the signature verification fails; otherwise, returns `nil`.

#### Description

The `VerifySignature` method checks the validity of the provided webhook signature against a computed HMAC signature generated from the payload. It performs the following:

1. Computes the HMAC signature using the payload and the API key.
2. Compares the computed signature with the received signature.
3. Returns an error if the signatures do not match, indicating an invalid signature; otherwise, it returns `nil`, indicating that the signature is valid.
