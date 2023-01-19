// Package models contain all structs, types, and constants
// that have been modelled for use in wallet-api.
package models

import "time"

// WalletTxn represents a transaction of the wallet.
type WalletTxn struct {
	DateTime time.Time `json:"datetime"`
	Amount   float32   `json:"amount"`
}
