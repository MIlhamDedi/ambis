package config

import (
	"fmt"
	"os"
)

var (
	AincradEndpoint    = os.Getenv("AINCRAD_ENDPOINT")
	KiritoEndpoint     = os.Getenv("KIRITO_ENDPOINT")
	HeathcliffEndpoint = os.Getenv("HEATHCLIFF_ENDPOINT")
	SterbenEndpoint    = os.Getenv("STERBEN_ENDPOINT")
	YuiEndpoint        = os.Getenv("YUI_ENDPOINT")
	AsunaEndpoint      = os.Getenv("ASUNA_ENDPOINT")
	SinonEndpoint      = os.Getenv("SINON_ENDPOINT")
)

func init() {
	// Log used configurations here
	fmt.Printf("Aincrad is listening on %s\n...", AincradEndpoint)
}
