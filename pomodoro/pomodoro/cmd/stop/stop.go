/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package stop

import (
	"fmt"

	"github.com/spf13/cobra"
)

func stop() {
	// get the started section and update the end time
	// save to the file with the name is ID
}

// stopCmd represents the stop command
var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a pomodoro section",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
