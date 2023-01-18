package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/udayangaac/wallet-api/models"
	"github.com/udayangaac/wallet-api/store/wallet"
	"github.com/udayangaac/wallet-api/timeconv"
)

func TestSummarizer_Save(t *testing.T) {

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	t1Converted := timeconv.GetNextHalfHour(t1)

	dummyErr := errors.New("dummy error")

	tests := []struct {
		MockTxn, Txn         models.WalletTxn
		MockErr, ExpectedErr error
	}{
		{
			MockTxn: models.WalletTxn{
				DateTime: t1Converted,
				Amount:   10.00,
			},
			MockErr: nil,
			Txn: models.WalletTxn{
				DateTime: t1,
				Amount:   10.00,
			},
			ExpectedErr: nil,
		},
		{
			MockTxn: models.WalletTxn{
				DateTime: t1Converted,
				Amount:   10.00,
			},
			MockErr: dummyErr,
			Txn: models.WalletTxn{
				DateTime: t1,
				Amount:   10.00,
			},
			ExpectedErr: dummyErr,
		},
		{
			MockTxn: models.WalletTxn{
				DateTime: t1Converted,
				Amount:   20.00,
			},
			MockErr: dummyErr,
			Txn: models.WalletTxn{
				DateTime: t1,
				Amount:   10.00,
			},
			ExpectedErr: wallet.ErrMismatch,
		},
	}

	for _, test := range tests {

		md := wallet.NewSummaryStoreMockData()
		md.AddToSaveOrUpdate(test.MockTxn, test.MockErr)

		ms := wallet.NewSummaryMockStore(md)
		s := NewSummarizer(ms)

		actualErr := s.Save(test.Txn)
		assert.Equal(t, test.ExpectedErr, actualErr)
	}

}

func TestSummarizer_GetHistory(t *testing.T) {

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	t1Converted := timeconv.GetNextHalfHour(t1)
	t2Converted := t1Converted.Add(60 * time.Minute)

	dummyErr := errors.New("dummy error")

	tests := []struct {
		To, From             time.Time
		ExpectedEntries      []models.WalletEntry
		MockEntries          []models.WalletEntry
		MockErr, ExpectedErr error
		MockParams           wallet.FilterParams
	}{
		{
			From: t1Converted,
			To:   t2Converted,
			MockParams: wallet.FilterParams{
				From: t1Converted,
				To:   t2Converted,
			},

			MockEntries: []models.WalletEntry{
				{
					DateTime: t1Converted,
					Balance:  20.00,
				},
				{
					DateTime: t2Converted,
					Balance:  35.00,
				},
			},

			ExpectedEntries: []models.WalletEntry{
				{
					DateTime: t1Converted,
					Balance:  35.00,
				},
			},
		}, {
			From: t1Converted,
			To:   t2Converted,
			MockParams: wallet.FilterParams{
				From: t1Converted,
				To:   t2Converted,
			},

			MockEntries: []models.WalletEntry{
				{
					DateTime: t1Converted,
					Balance:  20.00,
				},
			},

			ExpectedEntries: []models.WalletEntry{
				{
					DateTime: t1Converted,
					Balance:  20.00,
				},
			},
		}, {
			From: t1Converted,
			To:   t2Converted,
			MockParams: wallet.FilterParams{
				From: t1Converted,
				To:   t2Converted,
			},
			ExpectedErr: dummyErr,
			MockErr:     dummyErr,
		},
	}

	for _, test := range tests {

		md := wallet.NewSummaryStoreMockData()
		md.AddToGetAll(test.MockParams, test.MockEntries, test.MockErr)
		ms := wallet.NewSummaryMockStore(md)
		s := NewSummarizer(ms)

		actualEntries, actualErr := s.GetHistory(test.From, test.To)
		assert.Equal(t, test.ExpectedErr, actualErr)
		assert.Equal(t, test.ExpectedEntries, actualEntries)
	}
}

func TestSummarizer_GetLatestBalance(t *testing.T) {

	dummyErr := errors.New("dummy error")

	tests := []struct {
		MockEntry            models.WalletEntry
		MockErr, ExpectedErr error
		ExpectedBal          float32
	}{
		{
			MockEntry: models.WalletEntry{
				Balance: 10.00,
			},
			ExpectedBal: 10.00,
		},
		{
			MockErr:     dummyErr,
			ExpectedErr: dummyErr,
		},
	}

	for _, test := range tests {

		md := wallet.NewSummaryStoreMockData()
		md.AddToGetLast(test.MockEntry, test.MockErr)
		ms := wallet.NewSummaryMockStore(md)
		s := NewSummarizer(ms)

		actualBalance, actualErr := s.GetLatestBalance()
		assert.Equal(t, test.ExpectedErr, actualErr)
		assert.Equal(t, test.ExpectedBal, actualBalance)
	}
}
