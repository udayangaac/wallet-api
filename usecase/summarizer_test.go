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
		MockEntry, Entry     models.WalletEntry
		MockErr, ExpectedErr error
	}{
		{
			MockEntry: models.WalletEntry{
				DateTime: t1Converted,
				Balance:  10.00,
			},
			MockErr: nil,
			Entry: models.WalletEntry{
				DateTime: t1,
				Balance:  10.00,
			},
			ExpectedErr: nil,
		},
		{
			MockEntry: models.WalletEntry{
				DateTime: t1Converted,
				Balance:  10.00,
			},
			MockErr: dummyErr,
			Entry: models.WalletEntry{
				DateTime: t1,
				Balance:  10.00,
			},
			ExpectedErr: dummyErr,
		},
		{
			MockEntry: models.WalletEntry{
				DateTime: t1Converted,
				Balance:  20.00,
			},
			MockErr: dummyErr,
			Entry: models.WalletEntry{
				DateTime: t1,
				Balance:  10.00,
			},
			ExpectedErr: wallet.ErrMismatch,
		},
	}

	for _, test := range tests {

		md := wallet.NewSummaryStoreMockData()
		md.AddToSaveOrUpdate(test.MockEntry, test.MockErr)

		ms := wallet.NewSummaryMockStore(md)
		s := NewSummarizer(ms)

		actualErr := s.Save(test.Entry)
		assert.Equal(t, test.ExpectedErr, actualErr)
	}

}

func TestSummarizer_GetHistory(t *testing.T) {
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
			MockEntry: models.WalletEntry{
				Balance: 10.00,
			},
			ExpectedBal: 10.00,
			MockErr:     dummyErr,
			ExpectedErr: dummyErr,
		},
		{
			MockEntry: models.WalletEntry{
				Balance: 20.00,
			},
			ExpectedBal: 10.00,
			MockErr:     dummyErr,
			ExpectedErr: wallet.ErrMismatch,
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
