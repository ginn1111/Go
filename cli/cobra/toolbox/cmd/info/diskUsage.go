/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"
	"log"
	"strings"

	"os/exec"

	"github.com/spf13/cobra"
)

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "The usage memory of directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cd := exec.Command("du", "-sh")

		var out strings.Builder

		cd.Stdout = &out

		err := cd.Run()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Usage of current directory %v", out.String())

	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
