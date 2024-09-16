package cmd

import (
	"fmt"
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
			issuesResult, err := githubissues.GetListOfIssues(owner, repo)
			if err != nil {
				log.Fatalf("Error while getting issue: %v", err)
			} else {
				for _, issue := range issuesResult {
					fmt.Printf("#%-5d %9.20s %.55s \n\t%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
				}
			}
		} else if number == 0 {
			log.Println("You have to add a number of issue. Try Again")
		} else {
			issue, err := githubissues.GetIssue(owner, repo, number)
			if err != nil {
				log.Fatalf("Error while getting issue: %v", err)
			} else {
				fmt.Printf("#%-5d %9.20s %.55s \n\t%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
			}
		}

	},
}

func init() {
	getCmd.Flags().BoolVarP(&list, "list", "l", false, "List all issues in repo")
	getCmd.Flags().IntVarP(&number, "number", "n", 0, "Show issue with given number")
	rootCmd.AddCommand(getCmd)

}
