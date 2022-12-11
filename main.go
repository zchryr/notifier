package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func getResponseCodeInput() int {
	i, err := strconv.Atoi(os.Getenv("INPUT_RESPONSE_CODE"))

	if err != nil {
		panic(err)
	}

	return i
}

func main() {
	// Environment variables as input.
	body := os.Getenv("INPUT_BODY")
	url := os.Getenv("INPUT_URL")
	response_code := getResponseCodeInput()

	// Checking for verbose mode.
	verbose_value, _ := os.LookupEnv("INPUT_VERBOSE")
	if verbose_value == "true" {
		fmt.Println("Verbose output enabled.")
	}

	// HTTP client.
	client := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		panic(err)
	}

	// Send request.
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"User-Agent": {"https://github.com/zchryr/build-notifier-action"},
	}
	resp, err := client.Do(req)

	// Error handling.
	if resp.StatusCode != response_code {
		fmt.Println("Something is wrong :-(")
		panic(err)
	} else {
		if verbose_value == "true" {
			fmt.Println("Request sent successfully!")
		}
	}
}