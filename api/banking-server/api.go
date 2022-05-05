package bankingapi

import (
	"core-banking-server/api/banking-server/endpoint"
	bankingcoreconfig "core-banking-server/internal/configuration/banking-server"

	"github.com/gin-gonic/gin"
)

func SetupCreateSavingsAccountAPI(router *gin.Engine) {
	router.POST(bankingcoreconfig.DefaultConfig.CurrentVersion+"/savings/create", func(ctx *gin.Context) {
		endpoint.CreateSavingsAccount(ctx)
	})
}

func SetupSettleSavingsAccountAPI(router *gin.Engine) {
	router.POST(bankingcoreconfig.DefaultConfig.CurrentVersion+"/savings/settle", func(ctx *gin.Context) {
		endpoint.SettleSavingsAccount(ctx)
	})
}

func SetupFetchAccountInfoAPI(router *gin.Engine) {
	router.POST(bankingcoreconfig.DefaultConfig.CurrentVersion+"/account/info", func(ctx *gin.Context) {
		endpoint.FetchAccountInfo(ctx)
	})
}

func SetupLoginAPI(router *gin.Engine) {
	router.POST(bankingcoreconfig.DefaultConfig.CurrentVersion+"/account/login", func(ctx *gin.Context) {
		endpoint.Login(ctx)
	})
}

func SetupConfirmationAPI(router *gin.Engine) {
	router.POST(bankingcoreconfig.DefaultConfig.CurrentVersion+"/savings/confirmation", func(ctx *gin.Context) {
		endpoint.ConfirmTransaction(ctx)
	})
}
