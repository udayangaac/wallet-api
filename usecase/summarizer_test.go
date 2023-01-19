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

	from, _ := time.Parse(time.RFC3339, "2006-01-02T09:01:05+00:00")
	to, _ := time.Parse(time.RFC3339, "2006-01-02T11:00:05+00:00")

	fromMock, _ := time.Parse(time.RFC3339, "2006-01-02T09:30:00+00:00")
	toMock, _ := time.Parse(time.RFC3339, "2006-01-02T11:30:00+00:00")

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T10:00:00+00:00")
	t2, _ := time.Parse(time.RFC3339, "2006-01-02T10:30:00+00:00")

	t1Exp, _ := time.Parse(time.RFC3339, "2006-01-02T10:00:00Z")
	t2Exp, _ := time.Parse(time.RFC3339, "2006-01-02T11:00:00Z")

	dummyErr := errors.New("dummy error")

	tests := []struct {
		To, From             time.Time
		ExpectedEntries      []models.WalletEntry
		MockEntries          []models.WalletEntry
		MockErr, ExpectedErr error
		MockParams           wallet.FilterParams
	}{
		{
			From: from,
			To:   to,
			MockParams: wallet.FilterParams{
				From: fromMock,
				To:   toMock,
			},

			MockEntries: []models.WalletEntry{
				{
					DateTime: t1,
					Balance:  20.00,
				},
				{
					DateTime: t2,
					Balance:  35.00,
				},
			},

			ExpectedEntries: []models.WalletEntry{
				{
					DateTime: t1Exp,
					Balance:  20.00,
				},
				{
					DateTime: t2Exp,
					Balance:  35.00,
				},
			},
		},
		{
			From: from,
			To:   to,
			MockParams: wallet.FilterParams{
				From: fromMock,
				To:   toMock,
			},

			MockEntries: []models.WalletEntry{
				{
					DateTime: t1,
					Balance:  20.00,
				},
			},

			ExpectedEntries: []models.WalletEntry{
				{
					DateTime: t1Exp,
					Balance:  20.00,
				},
				{
					DateTime: t2Exp,
					Balance:  20.00,
				},
			},
		},
		{
			From: from,
			To:   to,
			MockParams: wallet.FilterParams{
				From: fromMock,
				To:   toMock,
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
