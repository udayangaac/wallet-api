// Package wallet includes wallet-related storage or repository implementations.
package wallet

import (
	"errors"

	"github.com/udayangaac/wallet-api/models"
)

var (
	ErrMismatch = errors.New("mismatch between expected and given values")
)

// MockDataSummaryStore contains dummy data for testing.
type SummaryStoreMockData struct {
	saveOrUpdate struct {
		Params struct {
			WalletEntry models.WalletEntry
		}
		Returns struct {
			Err error
		}
	}
	getAll struct {
		Params struct {
			FilterParams FilterParams
		}
		Returns struct {
			Entries []models.WalletEntry
			Err     error
		}
	}
	getLast struct {
		Returns struct {
			Entry models.WalletEntry
			Err   error
		}
	}
}

func NewSummaryStoreMockData() SummaryStoreMockData {
	return SummaryStoreMockData{}
}

func (m *SummaryStoreMockData) AddToSaveOrUpdate(entry models.WalletEntry, err error) {
	m.saveOrUpdate.Params.WalletEntry = entry
	m.saveOrUpdate.Returns.Err = err
}

func (m *SummaryStoreMockData) AddToGetLast(entry models.WalletEntry, err error) {
	m.getLast.Returns.Entry = entry
	m.getLast.Returns.Err = err
}

// NewSummaryMockStore create a new instance of mock implementation of SummaryStore.
func NewSummaryMockStore(data SummaryStoreMockData) SummaryStore {
	return &summaryMockPostgres{
		mockData: data,
	}
}

type summaryMockPostgres struct {
	mockData SummaryStoreMockData
}

// SaveOrUpdate saves a new entry or makes changes to an existing one. If there is no
// entry for the specified data and time, create a entry by adding the most recent balance
// to the provided balance. If not, update the current entry by adding the provided balance
// to the balance already present in that specific entry.
func (s *summaryMockPostgres) SaveOrUpdate(entry models.WalletEntry) (err error) {
	expected := s.mockData.saveOrUpdate.Params.WalletEntry
	if expected.Balance == entry.Balance && expected.DateTime.Equal(entry.DateTime) {
		return s.mockData.saveOrUpdate.Returns.Err
	}
	return ErrMismatch
}

// GetAll retrieves all wallet entries filtered given filters parameters.
func (s *summaryMockPostgres) GetAll(params FilterParams) (entries []models.WalletEntry, err error) {
	expected := s.mockData.getAll.Params.FilterParams
	if expected.From.Equal(params.From) && expected.To.Equal(params.To) {
		return s.mockData.getAll.Returns.Entries, s.mockData.getAll.Returns.Err
	}
	return []models.WalletEntry{}, ErrMismatch
}

// GetLast retrieve the last inserted wallet entry.
func (s *summaryMockPostgres) GetLast() (entry models.WalletEntry, err error) {
	return s.mockData.getLast.Returns.Entry, s.mockData.getLast.Returns.Err
}
