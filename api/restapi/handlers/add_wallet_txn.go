// Package models contain all http handlers.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udayangaac/wallet-api/api/restapi/operation"
	"github.com/udayangaac/wallet-api/usecase"
)

// GetAddWalletTxn returns a handler for save transaction to wallet.
func GetAddWalletTxn(s usecase.Summarizer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := operation.NewAddWalletTxnParams(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, operation.NewErrorResponse(err))
		}

		if err = s.Save(params.Body.ToWalletTxn()); err != nil {
			ctx.JSON(http.StatusInternalServerError, operation.NewErrorResponse(err))
		}

		ctx.JSON(http.StatusOK, operation.NewSuccessResponse("Save succeeded"))
	}
}
