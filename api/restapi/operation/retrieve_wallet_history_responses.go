package operation

import (
	"time"

	"github.com/udayangaac/wallet-api/models"
)

// NewRetrieveWalletBalanceResp create an instance of retrieve wallet balance response.
func NewRetrieveWalletHistoryResp(loc *time.Location, entries []models.WalletEntry) []models.WalletEntry {
	for i, _ := range entries {
		entries[i].DateTime = entries[i].DateTime.In(loc)
	}

	return entries
}
