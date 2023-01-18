package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udayangaac/wallet-api/api/restapi/operation"
	"github.com/udayangaac/wallet-api/usecase"
)

// RetrieveWalletHistory returns a handler function for retrieve wallet history.
func RetrieveWalletHistory(s usecase.Summarizer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := operation.NewRetrieveWalletHistoryParams(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, operation.NewErrorResponse(err))
		}

		entries, err := s.GetHistory(params.StartDatetime, params.EndDatetime)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, entries)
	}
}
