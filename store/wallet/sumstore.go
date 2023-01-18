// Package wallet includes wallet-related storage or repository implementations.
package wallet

import (
	"time"

	"github.com/udayangaac/wallet-api/models"
)

// FilterParams includes the filter parameters that are taken into account
// when obtaining the list of wallet entries.
type FilterParams struct {
	To   time.Time
	From time.Time
}

// SummaryStore is abstractions for wallet summaries implementations.
type SummaryStore interface {

	// SaveOrUpdate saves a new entry or makes changes to an existing one. If there is no
	// entry for the specified data and time, create a entry by adding the most recent balance
	// to the provided amount. If not, update the current entry by adding the provided amount
	// to the balance already present in that specific entry.
	SaveOrUpdate(txn models.WalletTxn) (err error)

	// GetAll retrieves all wallet entries filtered given filters parameters.
	GetAll(params FilterParams) (entries []models.WalletEntry, err error)

	// GetLast retrieve the last inserted wallet entry.
	GetLast() (entry models.WalletEntry, err error)
}
