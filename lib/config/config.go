package config

import (
	"fmt"
	"os"
)

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

func init() {
	// Log used configurations here
	fmt.Printf("Running %s Environment...\n", EnvironmentMode)
}
