/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package count

import (
	"fmt"

	"github.com/spf13/cobra"
)

// countCmd represents the count command
var CountCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the summary of sections and minutes in a day",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("count called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
