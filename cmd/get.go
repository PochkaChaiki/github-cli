/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	githubissues "github-cli/internal/githubIssues"
	"log"

	"github.com/spf13/cobra"
)

var list bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Issue(s)",
	Long:  `This command gets one issue by the number (-n flag) or list of issues (-l flag) and prints out.`,
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			err := githubissues.GetListOfIssues(owner, repo)
			if err != nil {
				log.Fatalf("Error while getting issue: %v", err)
			}
		} else if number == 0 {
			log.Fatalf("You have to add a number of issue. Try Again")
		} else {
			err := githubissues.GetIssue(owner, repo, number)
			if err != nil {
				log.Fatalf("Error while getting issue: %v", err)
			}
		}

	},
}

func init() {
	getCmd.Flags().BoolVarP(&list, "list", "l", false, "List all issues in repo")
	getCmd.Flags().IntVarP(&number, "number", "n", 0, "Show issue with given number")
	rootCmd.AddCommand(getCmd)

}
