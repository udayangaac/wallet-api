package operation

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/udayangaac/wallet-api/models"
)

// NewAddWalletTxnParams extract save wallet parameters from the request.
func NewAddWalletTxnParams(ctx *gin.Context) (AddWalletTxnParams, error) {
	body := AddWalletTxnBody{}
	err := ctx.BindJSON(&body)
	return AddWalletTxnParams{
		Body: &body,
	}, err
}

// AddWalletTxnParams request parameters of save wallet request.
type AddWalletTxnParams struct {
	Body *AddWalletTxnBody
}

// AddWalletTxnBody request body of add transaction to wallet request.
type AddWalletTxnBody struct {
	DateTime time.Time `json:"datetime"`
	Amount   float32   `json:"amount"`
}

// ToWalletTxn get models.ToWalletTxn from AddWalletTxnBody.
func (s AddWalletTxnBody) ToWalletTxn() models.WalletTxn {
	return models.WalletTxn{
		DateTime: s.DateTime,
		Amount:   s.Amount,
	}
}
