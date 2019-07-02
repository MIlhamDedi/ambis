package config

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Configurations passed by environment variable
var (
	// Running Environment
	EnvironmentMode = os.Getenv("ENV_MODE")

	// Database Configuration
	DBMode     = os.Getenv("DB_MODE")
	DBName     = os.Getenv("DB_NAME")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")

	// Internal Endpoint
	AincradEndpoint    = os.Getenv("AINCRAD_ENDPOINT")
	KiritoEndpoint     = os.Getenv("KIRITO_ENDPOINT")
	HeathcliffEndpoint = os.Getenv("HEATHCLIFF_ENDPOINT")
	SterbenEndpoint    = os.Getenv("STERBEN_ENDPOINT")
	YuiEndpoint        = os.Getenv("YUI_ENDPOINT")
	AsunaEndpoint      = os.Getenv("ASUNA_ENDPOINT")
	SinonEndpoint      = os.Getenv("SINON_ENDPOINT")

	// Received Root Path
	AincradRootPath    = os.Getenv("AINCRAD_ROOT_PATH")
	KiritoRootPath     = os.Getenv("KIRITO_ROOT_PATH")
	HeathcliffRootPath = os.Getenv("HEATHCLIFF_ROOT_PATH")
	SterbenRootPath    = os.Getenv("STERBEN_ROOT_PATH")
	YuiRootPath        = os.Getenv("YUI_ROOT_PATH")
	AsunaRootPath      = os.Getenv("ASUNA_ROOT_PATH")
	SinonRootPath      = os.Getenv("SINON_ROOT_PATH")

	// Secrets
	SigningSecret = os.Getenv("SIGNING_SECRET")
)

// Configurations passed by flag arguments
var (
	PortNumberPtr = flag.Int("p", 0, "port number to run the service on")
)

func init() {
	// Log used configurations here
	if EnvironmentMode == "" {
		log.Fatal("no valid configuration file has been sourced")
	}
	flag.Parse()
	fmt.Printf("Running in %s...\n", EnvironmentMode)
}

// GetPortAddr returns custom port from flag -f if specified,
// otherwise returns passed default port
func GetPortAddr(defaultAddr int) string {
	if *PortNumberPtr == 0 {
		return fmt.Sprintf(":%d", defaultAddr)
	}
	return fmt.Sprintf(":%d", *PortNumberPtr)
}

// SetLogMode sets Logrus configuration with pre-configured setup
func SetLogMode(envMode string) {
	// Setting Formatter
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Setting Log Output
	log.SetOutput(os.Stdout)

	// Setting Log Level
	switch envMode {
	case "PRODUCTION":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.TraceLevel)
	}
}
