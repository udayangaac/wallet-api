package operation

import (
	"time"

	"github.com/gin-gonic/gin"
)

// SaveWalletBody request body of save wallet request.
type RetrieveWalletHistoryParams struct {
	StartDatetime time.Time `json:"startDatetime"`
	EndDatetime   time.Time `json:"endDatetime"`
}

// NewRetrieveWalletHistoryParams extract retrieve wallet history parameter from the request.
func NewRetrieveWalletHistoryParams(ctx *gin.Context) (RetrieveWalletHistoryParams, error) {
	params := RetrieveWalletHistoryParams{}
	err := ctx.BindJSON(&params)
	return params, err
}
