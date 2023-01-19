// Package models contain all structs, types, and constants
// that have been modelled for use in wallet-api.
package models

import "errors"

var (
	// ErrInvalidDateTime an error which will be returned
	// when date time format is invalid.
	ErrInvalidDateTime = errors.New("invalid date time")

	// ErrInvalidAmount an error which will be returned
	// when amount is invalid.
	ErrInvalidAmount = errors.New("invalid amount")
)
