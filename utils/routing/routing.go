package routing

import (
	"bankserver/utils/api"

	"github.com/gin-gonic/gin"
)

func SetupCreateNewSavingsAccountAPI(router *gin.Engine) {
	router.POST("/v1/savings/create", func(ctx *gin.Context) {
		api.CreateNewSavingsAccount(ctx)
	})
}

func SetupFetchSavingsAccountAPI(router *gin.Engine) {
	router.POST("/v1/account/info", func(ctx *gin.Context) {
		api.FetchAccountInfo(ctx)
	})
}

func SetupLoginAPI(router *gin.Engine) {
	router.POST("/v1/account/login", func(ctx *gin.Context) {
		api.Login(ctx)
	})
}

func SetupSettleSavingsAccount(router *gin.Engine) {
	router.POST("/v1/savings/settle", func(ctx *gin.Context) {
		api.SettleSavingsAccout(ctx)

	})
}
