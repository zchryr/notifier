/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send data.",
	Long: `Send data to an endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		body, _ := cmd.Flags().GetString("body")
		url, _ := cmd.Flags().GetString("url")

		// Convert flag from string to int.
		response_flag, _ := cmd.Flags().GetString("response")
		response, err := strconv.Atoi(response_flag)

		if err != nil {
			panic(err)
		}

		send(body, url, response)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().String("body", "", "The data that should be sent in the body of the message.")
	sendCmd.PersistentFlags().String("url", "", "The URL where the data should be sent to.")
	sendCmd.PersistentFlags().String("response", "", "The URL where the data should be sent to.")

	sendCmd.MarkFlagsRequiredTogether("body", "url")
}

func send(body string, url string, response int) {
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
	if resp.StatusCode != response {
		fmt.Println("Server responded with status code: ", resp.StatusCode)
		panic(err)
	} else {
		fmt.Println("Request sent successfully!")
	}
}