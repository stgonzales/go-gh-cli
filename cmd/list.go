/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stgonzales/gh-cli/helpers"
)

type Repo struct {
	Name string `json:"name"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		req, reqErr := helpers.NewRequest("GET", "users/stgonzales/repos", nil)

		if reqErr != nil {
			log.Fatal(reqErr)
		}

		res, resErr := helpers.Client(req)

		if resErr != nil {
			log.Fatal(resErr)
		}

		repos := helpers.UnmarshallRepository(res)

		for _, repo := range repos {
			fmt.Printf("- %s\n", repo.Name)
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
