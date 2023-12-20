package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/siddharth1297/gocurl"
)

func main() {
	parser := argparse.NewParser("cccurl", "curl implementation in Go")

	url := parser.StringPositional(nil)
	version := parser.Flag("V", "version", &argparse.Options{Help: "Show version number and quit"})
	verbose := parser.Flag("v", "verbose", &argparse.Options{Help: "Make the operation more talkative"})
	method := parser.String("X", "method", &argparse.Options{Required: false, Help: "HTTP method"})
	headers := parser.StringList("H", "header", &argparse.Options{Required: false, Help: "HTTP headers"})
	payload := parser.String("d", "data", &argparse.Options{Required: false, Help: "POST Data"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *version {
		fmt.Println(gocurl.VERSION)
		return
	}

	curlConfig := gocurl.NewCurl()
	curlConfig.Url = *url
	curlConfig.Verbose = *verbose
	curlConfig.Method = *method
	curlConfig.Headers = *headers
	curlConfig.Payload = *payload

	if !curlConfig.VerifyCurlConfig() {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	curlConfig.StartCurl()
}
