package operation

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidStartDatetime an error which will be returned
	// when date time format is invalid.
	ErrInvalidStartDatetime = errors.New("invalid startDatetime")

	// ErrInvalidEndDatetime an error which will be returned
	// when date time format is invalid.
	ErrInvalidEndDatetime = errors.New("invalid endDatetime")
)

// RetrieveWalletHistoryParams request body of save wallet request.
type RetrieveWalletHistoryParams struct {
	StartDatetime time.Time `json:"startDatetime"`
	EndDatetime   time.Time `json:"endDatetime"`
}

// Validate validates the wallet history params.
func (a RetrieveWalletHistoryParams) Validate() error {
	if a.StartDatetime.IsZero() && a.EndDatetime.IsZero() {
		return fmt.Errorf("%s and %s", ErrInvalidStartDatetime, ErrInvalidEndDatetime)
	} else if a.StartDatetime.IsZero() {
		return ErrInvalidStartDatetime
	} else if a.EndDatetime.IsZero() {
		return ErrInvalidEndDatetime
	} else {
		return nil
	}
}

// NewRetrieveWalletHistoryParams extract retrieve wallet history parameter from the request.
func NewRetrieveWalletHistoryParams(ctx *gin.Context) (RetrieveWalletHistoryParams, error) {
	params := RetrieveWalletHistoryParams{}
	err := ctx.BindJSON(&params)
	if err != nil {
		return params, err
	}
	err = params.Validate()
	return params, err
}
