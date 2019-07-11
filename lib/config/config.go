package config

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Config serves standard App Configuration
type Config struct {
	// Runnning Application
	AppName AppIdentifier

	// Server Port Address
	PortAddr string

	// Running Environment
	EnvironmentMode EnvIdentifier

	// Database Configuration
	DBMode     string
	DBName     string
	DBUser     string
	DBPassword string

	// Internal Endpoint
	AincradEndpoint    string
	AlfheimEndpoint    string
	KiritoEndpoint     string
	HeathcliffEndpoint string
	SterbenEndpoint    string
	YuiEndpoint        string
	AsunaEndpoint      string
	SinonEndpoint      string

	// Received Root Path
	AincradRootPath string
	YuiRootPath     string
	AsunaRootPath   string
	SinonRootPath   string

	// Secrets
	SigningSecret string
}

// Indentifiers
type (
	// Application
	AppIdentifier string

	// Environment
	EnvIdentifier string
)

// Enums
const (
	// Applications
	AINCRAD    AppIdentifier = "aincrad"
	ALFHEIM    AppIdentifier = "alfheim"
	KIRITO     AppIdentifier = "kirito"
	HEATHCLIFF AppIdentifier = "heathcliff"
	STERBEN    AppIdentifier = "sterben"
	YUI        AppIdentifier = "yui"
	ASUNA      AppIdentifier = "asuna"
	SINON      AppIdentifier = "sinon"

	// Environments
	DEVELOPMENT EnvIdentifier = "development"
	STAGING     EnvIdentifier = "staging"
	PRODUCTION  EnvIdentifier = "production"
)

// Default Apps PortAddr
var defaultPortAddr = map[AppIdentifier]int{
	AINCRAD:    5000,
	ALFHEIM:    5001,
	KIRITO:     4000,
	HEATHCLIFF: 4001,
	STERBEN:    4002,
	YUI:        3000,
	ASUNA:      3001,
	SINON:      3002,
}

// Configurations passed by flag arguments
var (
	PortNumberPtr = flag.Int("p", 0, "port number to run the service on")
)

func init() {
	// Log used configurations here
	envMode := os.Getenv("ENV_MODE")
	if envMode == "" {
		log.Fatal("no valid configuration file has been sourced")
	}
	flag.Parse()
	fmt.Printf("Running in %s...\n", envMode)
}

// Configurations passed by environment variable
func Get(appEnum AppIdentifier) *Config {
	return &Config{
		AppName:            appEnum,
		PortAddr:           GetPortAddr(appEnum),
		EnvironmentMode:    EnvIdentifier(os.Getenv("ENV_MODE")),
		DBMode:             os.Getenv("DB_MODE"),
		DBName:             os.Getenv("DB_NAME"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		AincradEndpoint:    os.Getenv("AINCRAD_ENDPOINT"),
		KiritoEndpoint:     os.Getenv("KIRITO_ENDPOINT"),
		HeathcliffEndpoint: os.Getenv("HEATHCLIFF_ENDPOINT"),
		SterbenEndpoint:    os.Getenv("STERBEN_ENDPOINT"),
		YuiEndpoint:        os.Getenv("YUI_ENDPOINT"),
		AsunaEndpoint:      os.Getenv("ASUNA_ENDPOINT"),
		SinonEndpoint:      os.Getenv("SINON_ENDPOINT"),
		AincradRootPath:    os.Getenv("AINCRAD_ROOT_PATH"),
		YuiRootPath:        os.Getenv("YUI_ROOT_PATH"),
		AsunaRootPath:      os.Getenv("ASUNA_ROOT_PATH"),
		SinonRootPath:      os.Getenv("SINON_ROOT_PATH"),
		SigningSecret:      os.Getenv("SIGNING_SECRET"),
	}
}

// GetPortAddr returns custom port from flag -f if specified,
// otherwise returns passed default port
func GetPortAddr(appEnum AppIdentifier) string {
	if *PortNumberPtr == 0 {
		return fmt.Sprintf(":%d", defaultPortAddr[appEnum])
	}
	return fmt.Sprintf(":%d", *PortNumberPtr)
}
