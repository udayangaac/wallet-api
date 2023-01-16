// Package entity contain all database entities.
package entity

import (
	"time"

	"github.com/udayangaac/wallet-api/models"
	"gorm.io/gorm"
)

// Summary is a database entity for wallet summary.
// consist of necessary information related to the wallet summary.
type Summary struct {
	gorm.Model
	DateTime time.Time `gorm:"column:date_time;uniqueIndex"`
	Balance  float32   `gorm:"column:balance;type:FLOAT8"`
}

// ParseFromModel parses models.WalletEntry to Summary entity.
func (s *Summary) ParseFromModel(entry models.WalletEntry) {
	s.Balance = entry.Balance
	s.DateTime = entry.DateTime
}

// ToWalletEntry parses Summary entity to WalletEntry.
func (s *Summary) ToWalletEntry() models.WalletEntry {
	return models.WalletEntry{
		DateTime: s.DateTime,
		Balance:  s.Balance,
	}
}
