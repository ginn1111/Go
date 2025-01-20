/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go/gorm/bookstore/pkg/models"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all books",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:9010/book/", nil)

		resp, _ := http.DefaultClient.Do(req)

		fmt.Println(resp.StatusCode)

		if resp.StatusCode == http.StatusOK {
			var listBook []models.Book

			json.NewDecoder(resp.Body).Decode(&listBook)
			defer resp.Body.Close()

			for _, book := range listBook {
				fmt.Printf("ID: %v - Name %v, author %v - %v \n", book.ID, book.Name, book.Author, book.Publication)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
