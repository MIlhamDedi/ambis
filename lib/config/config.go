package config

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Configurations passed by environment variable
var (
	// Running Environment
	EnvironmentMode = os.Getenv("ENV_MODE")

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

func GetPortAddr(defaultAddr int) string {
	if *PortNumberPtr == 0 {
		return fmt.Sprintf(":%d", defaultAddr)
	}
	return fmt.Sprintf(":%d", *PortNumberPtr)
}
