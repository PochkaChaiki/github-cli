/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	createIssue "githubIssues/internal/createIssue"
	"log"

	"github.com/spf13/cobra"
)

var Title, Text string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Issue",
	Long:  `This command creates issue in owner's (flag -o) repo (flag -r)`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createIssue.CreateIssue(Owner, Repo, Title, Text)
		if err != nil {
			log.Fatalf("Error while creating issue: %v", err)
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&Title, "title", "t", "Title", "Title for creating issue")
	createCmd.Flags().StringVar(&Text, "text", "Initial text", "Body of the issue")
	rootCmd.AddCommand(createCmd)
}
