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
func (w *Summarizer) Save(txn models.WalletTxn) error {

	txn.DateTime = timeconv.GetNextHalfHour(txn.DateTime)

	return w.store.SaveOrUpdate(txn)
}

// GetHistory returns a history wallet balance at the end of each defined time periods between two date times.
func (w *Summarizer) GetHistory(from, to time.Time) ([]models.WalletEntry, error) {
	entries := make([]models.WalletEntry, 0)

	params := wallet.FilterParams{From: timeconv.GetNextHalfHour(from), To: timeconv.GetNextHalfHour(to)}

	dbEntries, err := w.store.GetAll(params)
	if err != nil {
		return nil, err
	}

	// create an temporarily map contains database records.
	tmpMap := make(map[string]models.WalletEntry)
	for _, entry := range dbEntries {
		tmpMap[timeconv.ToString(entry.DateTime)] = entry
	}

	nextHr := timeconv.GetNextHour(from)

	// Need to convert hover to same time zone.
	// Otherwise it is not possible to check the map with different time zones.
	nextHr = nextHr.UTC()

	// forwardBal is the balance which we need to
	// consider in the next hour which has no records.
	forwardBal := float32(0)

	entry, errLast := w.store.GetLast(nextHr)
	if errLast == nil {
		forwardBal = entry.Balance
	}

	n := int(to.Sub(from).Hours())
	for i := 0; i <= n; i++ {

		we := models.WalletEntry{
			DateTime: nextHr,
			Balance:  forwardBal,
		}

		entry, ok := tmpMap[timeconv.ToString(nextHr)]
		if ok {
			forwardBal = entry.Balance
			we.Balance = entry.Balance
		} else {
			preBin := timeconv.SubHalfHour(nextHr)
			entry, ok = tmpMap[timeconv.ToString(preBin)]
			if ok {
				forwardBal = entry.Balance
				we.Balance = entry.Balance
			}
		}

		entries = append(entries, we)
		nextHr = nextHr.Add(time.Hour * 1)
	}

	return entries, nil
}

// GetLatestBalance returns the latest balance.
func (w *Summarizer) GetLatestBalance() (float32, error) {
	entry, err := w.store.GetLast(time.Now())
	if err != nil {
		return 0, err
	}
	return entry.Balance, nil
}
