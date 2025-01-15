/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package test

import (
	"fmt"
	"log"
	"time"

	"github.com/go/pomodoro/pomodoro/models"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Pomodoro{
			ID:        123,
			Minute:    30,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Hour / 2),
		}

		// utils.Save(&po)
		// models.List()
		timer := models.Timer{
			CStop: make(chan bool, 1),
		}
		timer.Ticks()

		go func() {
			time.Sleep(time.Second * 3)
			log.Println("Stop")
			timer.Stops()
		}()

		for _ = range timer.CStop {
		}

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
