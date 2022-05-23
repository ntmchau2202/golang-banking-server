package bankingserverinit

import (
	"context"
	bankingapi "core-banking-server/api/banking-server"
	bankingcoreconfig "core-banking-server/internal/configuration/banking-server"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func InitConfig(envPath string) (err error) {
	bankingcoreconfig.DefaultConfig.DeployMode = os.Getenv("DEPLOY_MODE")
	if len(bankingcoreconfig.DefaultConfig.DeployMode) == 0 {
		if len(envPath) == 0 {
			envPath = ".env"
		}
		err = godotenv.Load(envPath)
		if err != nil {
			return
		}
	}

	bankingcoreconfig.DefaultConfig.CurrentVersion = os.Getenv("CURRENT_VERSION")
	bankingcoreconfig.DefaultConfig.Port = os.Getenv("DEFAULT_PORT")
	return
}

func InitEngine(mode string) (engine *gin.Engine) {
	gin.SetMode(mode)
	engine = gin.New()
	engine.Use(cors.Default())
	engine.Use(gin.Logger())
	return
}

func SetupAPIs(router *gin.Engine) {
	bankingapi.SetupCreateSavingsAccountAPI(router)
	bankingapi.SetupSettleSavingsAccountAPI(router)
	bankingapi.SetupFetchAccountInfoAPI(router)
	bankingapi.SetupLoginAPI(router)
	bankingapi.SetupConfirmationAPI(router)
	bankingapi.SetupRequestSignature(router)
}

func SetupGracefulShutdown(ctx context.Context, port string, engine *gin.Engine) {
	server := &http.Server{
		Addr:    ":" + port,
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
			log.Println("Application failed", err)
		}
	}()
	log.WithFields(log.Fields{"bind": port}).Info("Running application")

	stop := <-signalForExit
	log.Info("Stop signal Received", stop)
	log.Info("Waiting for all jobs to stop")
}
