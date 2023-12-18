package gocurl

import (
	"fmt"
	"os"
)

const VERSION = "curl/8.1.2"

type CurlConfig struct {
	Url     string
	Verbose bool
	Method  string
}

func NewCurl() *CurlConfig {
	return &CurlConfig{}
}

func printCurlError(code uint16) {
	fmt.Fprintf(os.Stderr, "curl: (%d) %s\n", code, errCodeToStr(code))
	os.Exit(int(code))
}
func (config *CurlConfig) VerifyCurlConfig() bool {
	correct := true
	if config.Url == "" {
		printCurlError(CURLE_URL_MALFORMAT)
		correct = false
	}
	// Don't apply any rule on Method
	return correct
}

func (config *CurlConfig) StartCurl() {
	config.step3()
}
