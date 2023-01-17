// Package usecase implements the main logic of wallet-api.
package usecase

import (
	"time"

	"github.com/udayangaac/wallet-api/models"
	"github.com/udayangaac/wallet-api/store/wallet"
	"github.com/udayangaac/wallet-api/timeconv"
)

// minPeriod minimum time difference between two timezones.
const minPeriod = 1800 * time.Second

// Summarizer summarize the wallet.
type Summarizer struct {
	store wallet.SummaryStore
}

// NewSummarizer create an instance of a Summarizer.
func NewSummarizer(store wallet.SummaryStore) Summarizer {
	return Summarizer{
		store: store,
	}
}

// Save saves a record.
func (w *Summarizer) Save(entry models.WalletEntry) error {

	entry.DateTime = timeconv.GetNextHalfHour(entry.DateTime)

	return w.store.SaveOrUpdate(entry)
}

// GetHistory returns a history wallet balance at the end of each defined time periods between two date times.
func (w *Summarizer) GetHistory(from, to time.Time) ([]models.WalletEntry, error) {

	from = timeconv.GetNextHalfHour(from)

	to = timeconv.GetNextHalfHour(to)

	entries := make([]models.WalletEntry, 0)

	params := wallet.FilterParams{
		From: from,
		To:   to,
	}

	dbEntries, err := w.store.GetAll(params)
	if err != nil {
		return nil, err
	}

	numOfEntryPerRecord := int(models.PeriodOfWalletHistory / minPeriod)

	for i := 0; i < len(dbEntries); {

		tmp := models.WalletEntry{}
		idx := i

		for j := 0; j < numOfEntryPerRecord; j++ {
			if i >= len(dbEntries) {
				break
			}

			if tmp.DateTime.IsZero() {
				tmp = dbEntries[idx+j]
			} else {
				tmp.Balance = dbEntries[idx+j].Balance
			}

			i++
		}

		entries = append(entries, tmp)
	}
	return entries, nil
}

// GetLatestBalance returns the latest balance.
func (w *Summarizer) GetLatestBalance() (float32, error) {
	entry, err := w.store.GetLast()
	if err != nil {
		return 0, err
	}
	return entry.Balance, nil
}
