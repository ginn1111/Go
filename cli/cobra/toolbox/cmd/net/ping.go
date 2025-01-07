/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

type Ping struct {
	Url string `validate:"hostname_rfc1123"`
}

func (p *Ping) ping() {
	if p.Url == "" {
		log.Fatal("hostname is required")
	}

	hostName := fmt.Sprintf("http://%s", p.Url)

	req, err := http.NewRequest(http.MethodHead, hostName, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(resp.StatusCode)
}

// pingCmd represents the ping command
var validate *validator.Validate
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the network",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ping := Ping{Url: cmd.Flag("url").Value.String()}

		validate = validator.New(validator.WithRequiredStructEnabled())
		// err := validate.Var(ping.url, "hostname")
		err := validate.Struct(&ping)

		if err != nil {
			log.Fatal(err)
			return
		}

		ping.ping()
	},
}

func init() {

	var urlFlag string
	pingCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "Host name")
	pingCmd.MarkFlagRequired("url")

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
