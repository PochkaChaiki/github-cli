/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	githubissues "github-cli/internal/githubIssues"
	"log"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Issue",
	Long:  `This command creates issue in owner's (flag -o) repo (flag -r)`,
	Run: func(cmd *cobra.Command, args []string) {
		err := githubissues.CreateIssue(owner, repo, title, text)
		if err != nil {
			log.Fatalf("Error while creating issue: %v", err)
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&title, "title", "t", "Title", "Title for creating issue")
	createCmd.Flags().StringVar(&text, "text", "Initial text", "Body of the issue")
	rootCmd.AddCommand(createCmd)
}
