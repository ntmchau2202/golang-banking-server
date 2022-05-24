package bankingcoreconfig

type Config struct {
	DeployMode     string
	Port           string
	CurrentVersion string
}

var DefaultConfig Config
