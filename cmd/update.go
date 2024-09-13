/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	githubissues "github-cli/internal/githubIssues"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Issue",
	Long:  `This command updates one issue by the number (-n flag).`,
	Run: func(cmd *cobra.Command, args []string) {
		githubissues.UpdateIssue(owner, repo, number, title, text)
	},
}

func init() {

	updateCmd.Flags().StringVarP(&title, "t", "title", "", "")
	rootCmd.AddCommand(updateCmd)
}
