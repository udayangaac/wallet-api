// Package wallet includes wallet-related storage or repository implementations.
package wallet

import (
	"errors"

	"github.com/udayangaac/wallet-api/models"
)

// MockDataSummaryStore contains dummy data for testing.
type MockDataSummaryStore struct {
	SaveOrUpdate struct {
		Params struct {
			WalletEntry models.WalletEntry
		}
		Returns struct {
			Err error
		}
	}
	GetAll struct {
		Params struct {
			FilterParams FilterParams
		}
		Returns struct {
			Entries []models.WalletEntry
			Err     error
		}
	}
	GetLast struct {
		Returns struct {
			Entry models.WalletEntry
			Err   error
		}
	}
}

// NewSummaryMockStore create a new instance of mock implementation of SummaryStore.
func NewSummaryMockStore(data MockDataSummaryStore) SummaryStore {
	return &summaryMockPostgres{
		mockData: data,
	}
}

type summaryMockPostgres struct {
	mockData MockDataSummaryStore
}

// SaveOrUpdate saves a new entry or makes changes to an existing one. If there is no
// entry for the specified data and time, create a entry by adding the most recent balance
// to the provided balance. If not, update the current entry by adding the provided balance
// to the balance already present in that specific entry.
func (s *summaryMockPostgres) SaveOrUpdate(entry models.WalletEntry) (err error) {
	expected := s.mockData.SaveOrUpdate.Params.WalletEntry
	if expected == entry {
		return s.mockData.GetAll.Returns.Err
	}
	return errors.New("given WalletEntry is not equal to expected WalletEntry")
}

// GetAll retrieves all wallet entries filtered given filters parameters.
func (s *summaryMockPostgres) GetAll(params FilterParams) (entries []models.WalletEntry, err error) {
	expected := s.mockData.GetAll.Params.FilterParams
	if expected == params {
		return s.mockData.GetAll.Returns.Entries, s.mockData.GetAll.Returns.Err
	}
	return []models.WalletEntry{}, errors.New("given FilterParams are not equal to expected FilterParams")
}

// GetLast retrieve the last inserted wallet entry.
func (s *summaryMockPostgres) GetLast() (entry models.WalletEntry, err error) {
	return s.mockData.GetLast.Returns.Entry, s.mockData.GetLast.Returns.Err
}
