// Package handler contains all http handlers.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udayangaac/wallet-api/api/restapi/operation"
	"github.com/udayangaac/wallet-api/usecase"
)

// GetRetrieveWalletHistory returns a handler function for retrieve wallet history.
func GetRetrieveWalletHistory(s usecase.Summarizer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := operation.NewRetrieveWalletHistoryParams(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, operation.NewErrorResponse(err))
			return
		}
		entries, err := s.GetHistory(params.StartDatetime, params.EndDatetime)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		resp := operation.NewRetrieveWalletHistoryResp(params.EndDatetime.Location(), entries)
		ctx.JSON(http.StatusOK, resp)
	}
}
