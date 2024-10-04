package chargily

//Chargily base Api endpoints
const (
	// The Chargily API endpoint for testing while developing
	TestAPIBaseUrl = "https://pay.chargily.net/test/api/v2/"

	// The Chargily API endpoint for production
    ProdAPIBaseUrl = "https://pay.chargily.net/api/v2/"
)



//Modes to be used in this SDK to interact with the api of Chargily
//Custom type for Mode
type Mode string

//Allowed modes as constants
const (
    Prod Mode = "prod"
    Test Mode = "test"
)