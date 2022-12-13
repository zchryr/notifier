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
		repo, _ := cmd.Flags().GetString("repo")
		workflow, _ := cmd.Flags().GetString("workflow")

		send(body, url, response, repo, workflow)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().String("body", "", "The data that should be sent in the body of the message.")
	sendCmd.PersistentFlags().String("url", "", "The URL where the data should be sent to.")
	sendCmd.PersistentFlags().String("response", "", "The expected successful response code from the server.")
	sendCmd.PersistentFlags().String("repo", "", "Name of the GitHub repo this CLI is used in.")
	sendCmd.PersistentFlags().String("workflow", "", "Name of the GitHub workflow this CLI is used in.")

	sendCmd.MarkFlagsRequiredTogether("body", "url", "response", "repo", "workflow")
}

func send(body string, url string, response string, repo string, workflow string) {
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
		"repo": {repo},
		"workflow": {workflow},
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