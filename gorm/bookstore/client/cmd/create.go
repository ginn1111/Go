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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new book",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, author, publication := cmd.Flag("name").Value.String(), cmd.Flag("author").Value.String(), cmd.Flag("publication").Value.String()

		book := models.Book{
			Name:        name,
			Author:      author,
			Publication: publication,
		}

		mBook, _ := json.Marshal(book)

		reader := bytes.NewReader(mBook)

		request, _ := http.NewRequest(http.MethodPost, "http://localhost:9010/book/", reader)

		resp, _ := http.DefaultClient.Do(request)

		fmt.Println(resp.StatusCode)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	var name, author, publication string

	createCmd.Flags().StringVarP(&name, "name", "n", "", "name of book")
	createCmd.Flags().StringVarP(&author, "author", "a", "", "author of book")
	createCmd.Flags().StringVarP(&publication, "publication", "b", "", "publication of book")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("author")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
