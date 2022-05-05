package main

import (
	"context"
	bankingcoreconfig "core-banking-server/internal/configuration/banking-server"
	bankingserverinit "core-banking-server/internal/initialize/banking-server"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	engine *gin.Engine
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	err := bankingserverinit.InitConfig("")
	if err != nil {
		log.Panic(err)
	}
	engine = bankingserverinit.InitEngine(bankingcoreconfig.DefaultConfig.DeployMode)
}

func main() {
	bankingserverinit.SetupAPIs(engine)
	bankingserverinit.SetupGracefulShutdown(ctx, bankingcoreconfig.DefaultConfig.Port, engine)

}
