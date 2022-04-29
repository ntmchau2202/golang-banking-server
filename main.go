package main

import (
	"bankserver/database"
	"bankserver/utils/routing"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	engine *gin.Engine
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())

	gin.SetMode("test")
	engine = gin.New()

	engine.Use(cors.Default())
	engine.Use(gin.Logger())
	database.GetDBConnection()
}

func main() {
	routing.SetupCreateNewSavingsAccountAPI(engine)
	routing.SetupFetchSavingsAccountAPI(engine)
	routing.SetupLoginAPI(engine)
	routing.SetupSettleSavingsAccount(engine)

	// setup gracefull shutdown

	server := &http.Server{
		Addr:    ":9999",
		Handler: engine,
	}
	signalForExit := make(chan os.Signal, 1)
	signal.Notify(signalForExit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Application failed", err)
		}
	}()
	log.WithFields(log.Fields{"bind": "9999"}).Info("Running application")

	stop := <-signalForExit
	log.Info("Stop signal Received", stop)
	log.Info("Waiting for all jobs to stop")
}
