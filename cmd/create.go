package cmd

import (
	"fmt"
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
		status, err := githubissues.CreateIssue(owner, repo, title, text, assignee, milestone, assignees)
		if err != nil {
			log.Fatalf("Error while creating issue: %v", err)
		} else {
			fmt.Printf("Status: %d\n", status)
		}

	},
}

func init() {
	createCmd.Flags().StringVarP(&title, "title", "t", "Title", "Title for creating issue")
	createCmd.Flags().StringVar(&text, "text", "Initial text", "Body of the issue")
	createCmd.Flags().StringVarP(&assignee, "assignee", "a", "", "The assignee to an issue")
	createCmd.Flags().Var(&assignees, "assignees", "A list of assignees to assign to an issue")
	createCmd.Flags().StringVarP(&milestone, "milestone", "m", "", "A milestone of an issue")
	rootCmd.AddCommand(createCmd)
}
