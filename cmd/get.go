/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	getissue "githubIssues/internal/getIssue"
	"log"

	"github.com/spf13/cobra"
)

var Number int
var List bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Issue(s)",
	Long:  `This command gets one issue by the number (-n flag) or list of issues (-l flag) and prints out.`,
	Run: func(cmd *cobra.Command, args []string) {
		if List {
			err := getissue.GetListOfIssues(Owner, Repo)
			if err != nil {
				log.Fatalf("Error while getting issue: %v", err)
			}
		} else if Number == 0 {
			log.Fatalf("You have to add a number of issue. Try Again")
		} else {
			err := getissue.GetIssue(Owner, Repo, Number)
			if err != nil {
				log.Fatalf("Error while getting issue: %v", err)
			}
		}

	},
}

func init() {
	getCmd.Flags().BoolVarP(&List, "list", "l", false, "List all issues in repo")
	getCmd.Flags().IntVarP(&Number, "number", "n", 0, "Show issue with given number")
	rootCmd.AddCommand(getCmd)

}
