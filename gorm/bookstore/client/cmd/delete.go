/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a book",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id := cmd.Flag("id").Value.String()
		req, _ := http.NewRequest(http.MethodDelete, "http://localhost:9010/book/"+id, nil)

		resp, _ := http.DefaultClient.Do(req)

		fmt.Println(resp.StatusCode)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	var id string
	deleteCmd.Flags().StringVarP(&id, "id", "i", "", "id of book")

	deleteCmd.MarkFlagRequired("id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
