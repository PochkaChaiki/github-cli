package cmd

import (
	"fmt"
	githubissues "github-cli/internal/githubIssues"
	"log"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Issue",
	Long:  `This command updates one issue by the number (-n flag).`,
	Run: func(cmd *cobra.Command, args []string) {
		if number == 0 {
			fmt.Println("You have to add a number of issue. Try Again")
		}
		status, err := githubissues.UpdateIssue(owner, repo, number, title, text, assignee, assignees, state, stateReason, milestone)
		if err != nil {
			log.Fatalf("Error while creating issue: %v", err)
		} else {
			fmt.Printf("Status: %d\n", status)
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "The title of an issue")
	updateCmd.Flags().IntVarP(&number, "number", "n", 0, "The title of an issue")
	updateCmd.Flags().StringVar(&text, "text", "", "Text that describes an issue")
	updateCmd.Flags().StringVarP(&assignee, "assignee", "a", "", "The assignee to an issue")
	updateCmd.Flags().Var(&assignees, "assignees", "A list of assignees assigned to an issue")
	updateCmd.Flags().StringVarP(&state, "state", "s", "", "A state of an issue")
	updateCmd.Flags().StringVar(&stateReason, "stateReason", "", "A state reason of an issue")
	updateCmd.Flags().StringVarP(&milestone, "milestone", "m", "", "A milestone of an issue")

	rootCmd.AddCommand(updateCmd)
}
