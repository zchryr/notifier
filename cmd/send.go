/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send data.",
	Long: `Send data to an endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(cmd.Flags().Count())
		// TODO not handling when no args are passed, somehow.

		body, _ := cmd.Flags().GetString("body")
		url, _ := cmd.Flags().GetString("url")
		response, _ := cmd.Flags().GetString("response")

		send(body, url, response)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().String("body", "", "The data that should be sent in the body of the message.")
	sendCmd.PersistentFlags().String("url", "", "The URL where the data should be sent to.")
	sendCmd.PersistentFlags().String("response", "", "The expected successful response code from the server.")

	sendCmd.MarkFlagsRequiredTogether("body", "url", "response")
}

func send(body string, url string, response string) {
	response_code, err := strconv.Atoi(response)

	if err != nil {
		panic(err)
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
		fmt.Println("Server responded with status code:", resp.StatusCode)
		fmt.Println("Expected response code:", response_code)
		os.Exit(1)
	} else if resp.StatusCode == response_code {
		fmt.Println("Request sent successfully!")
	} else if err != nil {
		panic(err)
	}
}