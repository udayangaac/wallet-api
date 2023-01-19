package restapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/udayangaac/wallet-api/api/restapi/handler"
	"github.com/udayangaac/wallet-api/dbconn"
	"github.com/udayangaac/wallet-api/store/entity"
	"github.com/udayangaac/wallet-api/store/wallet"
	"github.com/udayangaac/wallet-api/usecase"
)

// ConfigureAPI configure rest api.
func ConfigureAPI(port int) error {

	dbConn, err := dbconn.NewPGConnecter()
	if err != nil {
		return err
	}

	dbConn.AddEntities(&entity.Summary{})

	if err != nil {
		panic(err)
	}

	ss := wallet.NewSummaryPgStore(dbConn.GetDB())
	s := usecase.NewSummarizer(ss)

	r := gin.Default()

	v1 := r.Group("/v1")
	walGp := v1.Group("/wallet")

	walGp.POST("/transaction", handler.GetAddWalletTxn(s))
	walGp.GET("/history", handler.GetRetrieveWalletHistory(s))
	walGp.GET("/balance", handler.GetRetrieveWalletBalance(s))

	go r.Run(fmt.Sprintf(":%v", port))

	return nil
}
