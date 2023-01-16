// Package models contain all structs, types, and constants
// that have been modelled for use in wallet-api.
package models

import "time"

// WalletEntry represents an aggregated transactions in a defined time period.
type WalletEntry struct {
	DateTime time.Time `json:"datetime"`
	Balance  float32   `json:"amount"`
}
