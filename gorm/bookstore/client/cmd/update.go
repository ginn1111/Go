/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go/gorm/bookstore/pkg/models"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a book",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		book := models.Book{
			Author:      cmd.Flag("author").Value.String(),
			Name:        cmd.Flag("name").Value.String(),
			Publication: cmd.Flag("publication").Value.String(),
		}

		bookId := cmd.Flag("id").Value.String()

		mBook, _ := json.Marshal(book)

		req, _ := http.NewRequest(http.MethodPut, "http://localhost:9010/book/"+bookId, bytes.NewReader(mBook))

		resp, _ := http.DefaultClient.Do(req)

		fmt.Println(resp.StatusCode)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	var id, author, name, publication string

	updateCmd.Flags().StringVarP(&id, "id", "i", "", "id of book")
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "name of book")
	updateCmd.Flags().StringVarP(&author, "author", "a", "", "author of book")
	updateCmd.Flags().StringVarP(&publication, "publication", "b", "", "publication of book")

	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagsOneRequired("name", "author", "publication")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
