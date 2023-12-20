package gocurl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func printHttpHeader(headers http.Header, direction string) {
	for key, value := range headers {
		if direction == "" {
			fmt.Printf("%s: %s\n", key, value[0])
		} else {
			fmt.Printf("%s %s: %s\n", direction, key, value[0])
		}
	}
}

func handleError(err error, funct string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", funct, err.Error())
		os.Exit(1)
	}
}

func step1(url string) {
	var client http.Client // Underlying transport info
	var req *http.Request  // Http request details

	req, err := http.NewRequest("GET", url, nil)
	handleError(err, "NewRequest")

	// Add Headers
	req.Header.Add("Host", req.URL.Hostname())
	req.Header.Add("Accept", "*/*")

	fmt.Println("Connecting to ", req.URL.Hostname())
	fmt.Println("Sending request ", req.Method, req.URL.Path, req.Proto)
	printHttpHeader(req.Header, "")

	resp, err := client.Do(req)
	handleError(err, "client.Do")

	resp.Body.Close()
}

func step2(url string) {
	var client http.Client // Underlying transport info
	var req *http.Request  // Http request details

	req, err := http.NewRequest("GET", url, nil)
	handleError(err, "NewRequest")

	// Add Headers
	req.Header.Add("Host", req.URL.Hostname())
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "close")

	fmt.Println("Connecting to ", req.URL.Hostname())
	fmt.Println("Sending request ", req.Method, req.URL.Path, req.Proto)
	printHttpHeader(req.Header, "")

	resp, err := client.Do(req)
	handleError(err, "client.Do")

	defer resp.Body.Close()

	fmt.Println()

	fmt.Println(resp.Proto, resp.Status)

	printHttpHeader(resp.Header, "")

	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err, "ReadAll")

	fmt.Printf("%s", body)
}

func (config *CurlConfig) step3() {
	var client http.Client // Underlying transport info
	var req *http.Request  // Http request details

	req, err := http.NewRequest(config.Method, config.Url, nil)
	handleError(err, "NewRequest")

	// Add Headers
	req.Header.Add("Host", req.URL.Hostname())
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", VERSION)
	req.Header.Add("Connection", "close")

	if config.Verbose {
		//fmt.Println("Connecting to ", req.URL.Hostname())
		//fmt.Println("Sending request ", req.Method, " ", req.URL.Path, " ", req.Proto)
		fmt.Println(">", req.Method, req.URL.Path, req.Proto)
		printHttpHeader(req.Header, ">")
	}
	resp, err := client.Do(req)
	handleError(err, "client.Do")

	defer resp.Body.Close()

	if config.Verbose {

		fmt.Println("> ")

		fmt.Println("<", resp.Proto, resp.Status)

		printHttpHeader(resp.Header, "<")

		fmt.Println("<")
	}
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err, "ReadAll")

	fmt.Printf("%s", body)
}

func (config *CurlConfig) step4() {
	var client http.Client // Underlying transport info
	var req *http.Request  // Http request details

	req, err := http.NewRequest(config.Method, config.Url, bytes.NewBuffer([]byte(config.Payload)))
	handleError(err, "NewRequest")

	// Add Headers
	req.Header.Add("Host", req.URL.Hostname())
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", VERSION)
	req.Header.Add("Connection", "close")

	if req.Method == "POST" {
		for _, header := range config.Headers {
			kv_pair := strings.Split(header, ":")
			req.Header.Add(kv_pair[0], kv_pair[1])
		}
		req.Header.Add("Content-Length", strconv.Itoa((len(config.Payload))))
	}

	if config.Verbose {
		//fmt.Println("Connecting to ", req.URL.Hostname())
		//fmt.Println("Sending request ", req.Method, " ", req.URL.Path, " ", req.Proto)
		fmt.Println(">", req.Method, req.URL.Path, req.Proto)
		printHttpHeader(req.Header, ">")
	}
	resp, err := client.Do(req)
	handleError(err, "client.Do")

	defer resp.Body.Close()

	if config.Verbose {

		fmt.Println("> ")

		fmt.Println("<", resp.Proto, resp.Status)

		printHttpHeader(resp.Header, "<")

		fmt.Println("<")
	}
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err, "ReadAll")

	fmt.Printf("%s", body)
}

func (config *CurlConfig) step5() {
	var client http.Client // Underlying transport info
	var req *http.Request  // Http request details

	req, err := http.NewRequest(config.Method, config.Url, bytes.NewBuffer([]byte(config.Payload)))
	handleError(err, "NewRequest")

	// Add Headers
	req.Header.Add("Host", req.URL.Hostname())
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", VERSION)

	if req.Method == "POST" || req.Method == "PUT" {
		for _, header := range config.Headers {
			kv_pair := strings.Split(header, ":")
			req.Header.Add(kv_pair[0], kv_pair[1])
		}
		req.Header.Add("Content-Length", strconv.Itoa((len(config.Payload))))
	}

	if config.Verbose {
		//fmt.Println("Connecting to ", req.URL.Hostname())
		//fmt.Println("Sending request ", req.Method, " ", req.URL.Path, " ", req.Proto)
		fmt.Println(">", req.Method, req.URL.Path, req.Proto)
		printHttpHeader(req.Header, ">")
	}
	resp, err := client.Do(req)
	handleError(err, "client.Do")

	defer resp.Body.Close()

	if config.Verbose {

		fmt.Println("> ")

		fmt.Println("<", resp.Proto, resp.Status)

		printHttpHeader(resp.Header, "<")

		fmt.Println("<")
	}
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err, "ReadAll")

	fmt.Printf("%s", body)
}
