// Package wallet includes wallet-related storage or repository implementations.
package wallet

import (
	"time"

	"github.com/udayangaac/wallet-api/models"
	entity "github.com/udayangaac/wallet-api/store/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NewSummaryPgStore create a new instance of Postgres implementation of SummaryStore.
func NewSummaryPgStore(db gorm.DB) SummaryStore {
	return &summaryPgPostgres{
		DB: db,
	}
}

type summaryPgPostgres struct {
	DB gorm.DB
}

// SaveOrUpdate saves a new entry or makes changes to an existing one. If there is no
// entry for the specified data and time, create a entry by adding the most recent balance
// to the provided balance. If not, update the current entry by adding the provided balance
// to the balance already present in that specific entry.
func (s *summaryPgPostgres) SaveOrUpdate(entry models.WalletEntry) (err error) {

	// Set the balance for existing row.
	rowBal := clause.OnConflict{
		Columns:   []clause.Column{{Name: "date_time"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"balance": gorm.Expr("summaries.balance + ?", entry.Balance)}),
	}

	// Set the balance for new row.
	newRowBal := clause.Expr{
		SQL:  " ? + (SELECT balance FROM summaries WHERE summaries.deleted_at IS NULL ORDER BY summaries.id DESC LIMIT 1)",
		Vars: []interface{}{entry.Balance}}

	values := map[string]interface{}{
		"date_time":  entry.DateTime,
		"balance":    newRowBal,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	tx := s.DB.Model(entity.Summary{})

	return tx.Clauses(rowBal).Create(values).Error
}

// GetAll retrieves all wallet entries filtered given filters parameters.
func (s *summaryPgPostgres) GetAll(params FilterParams) (entries []models.WalletEntry, err error) {
	var tx *gorm.DB
	ses := make([]entity.Summary, 0)

	tx = s.DB.Model(entity.Summary{})

	if !params.From.IsZero() {
		tx = s.DB.Where("date_time >= ?", params.From)
	}

	if !params.To.IsZero() {
		tx = s.DB.Where("date_time < ?", params.To)
	}

	if !params.To.IsZero() && !params.From.IsZero() {
		tx = s.DB.Where("date_time >= ? AND date_time <= ?", params.From, params.To)
	}

	err = tx.Find(&ses).Error
	if err != nil {
		return
	}

	for _, ent := range ses {
		entries = append(entries, ent.ToWalletEntry())
	}

	return
}

// GetLast retrieve the last inserted wallet entry.
func (s *summaryPgPostgres) GetLast() (entry models.WalletEntry, err error) {
	se := entity.Summary{}

	err = s.DB.Last(&se).Error
	if err != nil {
		return
	}

	entry = se.ToWalletEntry()
	return
}
