/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package test

import (
	"fmt"
	"time"

	"github.com/go/pomodoro/pomodoro/models"
	"github.com/go/pomodoro/pomodoro/utils"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		po := models.Pomodoro{
			ID:        123,
			Minute:    30,
			StartTime: time.Now(),
			EndTime:   time.Now(),
		}

		utils.Save(&po)
		models.List()
		fmt.Println("test called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
