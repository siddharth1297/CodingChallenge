package gocurl

import (
	"fmt"
	"os"
)

// https://curl.se/libcurl/c/libcurl-errors.html

const CURLE_OK = 0
const CURLE_UNSUPPORTED_PROTOCOL = 1
const CURLE_FAILED_INIT = 2
const CURLE_URL_MALFORMAT = 3
const CURLE_UNKNOWN_ERROR = 4

func errCodeToStr(code uint16) string {
	switch code {
	case CURLE_OK:
		{
			return "CURLE_OK"
		}
	case CURLE_UNSUPPORTED_PROTOCOL:
		{
			return "CURLE_UNSUPPORTED_PROTOCOL"
		}
	case CURLE_FAILED_INIT:
		{
			return "CURLE_FAILED_INIT"
		}
	case CURLE_URL_MALFORMAT:
		{
			return "CURLE_URL_MALFORMAT"
		}
	case CURLE_UNKNOWN_ERROR:
		{
			return "CURLE_UNKNOWN_ERROR"
		}
	}
	fmt.Fprintf(os.Stderr, "%d: %s", code, "Unnamed Error")
	os.Exit(CURLE_UNKNOWN_ERROR)
	return ""
}
