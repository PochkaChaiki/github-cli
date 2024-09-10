/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	createIssue "githubIssues/internal"
	"log"

	"github.com/spf13/cobra"
)

var Owner, Repo, Title, Text string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Issue",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createIssue.CreateIssue(Owner, Repo, Title, Text)
		if err != nil {
			log.Fatalf("Error while creating issue: %v", err)
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&Owner, "owner", "o", "", "Owner's name of the repo to create issue in one")
	createCmd.Flags().StringVarP(&Repo, "repo", "r", "", "Repo's name to create issue in one")
	createCmd.Flags().StringVarP(&Title, "title", "t", "Title", "Title for creating issue")
	createCmd.Flags().StringVar(&Text, "text", "Initial text", "Body of the issue")
	rootCmd.AddCommand(createCmd)
}
