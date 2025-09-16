package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// Example function to verify the signature
func VerifySig( payload []byte, signature string , client * chargily.Client){
	// pass the payload and the signature into the client verifier
	err := client.Webhook.VerifySignature(payload, signature)

	// Handle errors
	if err != nil {
		fmt.Printf("Error verifying signature %v\n", err)
	}
}