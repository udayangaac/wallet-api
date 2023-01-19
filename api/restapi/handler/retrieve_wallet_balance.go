// Package handler contains all http handlers.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udayangaac/wallet-api/api/restapi/operation"
	"github.com/udayangaac/wallet-api/usecase"
)

// GetRetrieveWalletBalance returns a handler function for retrieve wallet balance.
func GetRetrieveWalletBalance(s usecase.Summarizer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bal, err := s.GetLatestBalance()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, operation.NewRetrieveWalletBalanceResp(bal))
	}
}
