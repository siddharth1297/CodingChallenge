package gocurl

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const VERSION = "curl/http/5"

type CurlConfig struct {
	Url     string
	Verbose bool
	Method  string
	Headers []string
	Payload string
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
	switch versions := strings.Split(VERSION, "/"); versions[1] {
	case "http":
		{
			switch step, _ := strconv.Atoi(versions[2]); step {
			case 1:
				step1(config.Url)
			case 2:
				step2(config.Url)
			case 3:
				config.step3()
			case 4:
				config.step4()
			case 5:
				config.step5()
			default:
				panic("Invalid step number in http")
			}
		}
	case "tcp":
		{
			panic("Not Implemented tcp")
		}
	}
}
