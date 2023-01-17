// Package models contain all structs, types, and constants
// that have been modelled for use in wallet-api.
package models

import "time"

// DateTimeFormat a domain specific date time format.
const DateTimeFormat = "2006-01-02T15:04:05-07:00"

// PeriodOfWalletHistory time difference between two wallet history records.
const PeriodOfWalletHistory = 3600 * time.Second
