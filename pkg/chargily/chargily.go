package chargily

// Wallet represents the balance details for a specific currency.
type Wallet struct {
    Currency        string `json:"currency"`         // The currency of the wallet (e.g., "dzd", "usd", "eur")
    Balance         int64  `json:"balance"`          // The total balance in the wallet
    ReadyForPayout  string `json:"ready_for_payout"` // The amount ready for payout
    OnHold          int64  `json:"on_hold"`         // The amount on hold
}

// Balance represents the overall balance entity.
type Balance struct {
    Entity   string   `json:"entity"`   // The entity type (e.g., "balance")
    LiveMode bool     `json:"livemode"` // Indicates whether the mode is live
    Wallets  []Wallet `json:"wallets"`  // A slice of Wallet structs
}
