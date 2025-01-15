/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

func startASection() {
	// create new pomodoro section
	// start tick timer
}

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro section with the optional minutes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
