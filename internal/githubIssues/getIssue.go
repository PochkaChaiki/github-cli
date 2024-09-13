package githubissues

import (
	"encoding/json"
	"fmt"
	getenvvars "github-cli/internal/getEnvVars"
	"net/http"
	"strings"
)

func GetListOfIssues(owner string, repo string) error {

	req, err := http.NewRequest(http.MethodGet, strings.Join([]string{getenvvars.GitHubAPI, owner, repo, "issues"}, "/"), nil)
	if err != nil {
		return fmt.Errorf("internal error: failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Beaver "+getenvvars.GitHubToken)
	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("internal error: failed to send request: %v", err)
	}
	var IssuesResult []Issue
	err = json.NewDecoder(resp.Body).Decode(&IssuesResult)
	if err != nil {
		return fmt.Errorf("internal error: failed to decode json: %v", err)
	}
	for _, issue := range IssuesResult {
		fmt.Printf("#%-5d %9.9s %.55s \n\t%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
	}
	return nil
}

func GetIssue(owner string, repo string, number int) error {

	req, err := http.NewRequest(http.MethodGet, strings.Join([]string{getenvvars.GitHubAPI, owner, repo, "issues", fmt.Sprint(number)}, "/"), nil)
	if err != nil {
		return fmt.Errorf("internal error: failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Beaver "+getenvvars.GitHubToken)
	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("internal error: failed to send request: %v", err)
	}
	var Issue Issue
	err = json.NewDecoder(resp.Body).Decode(&Issue)
	if err != nil {
		return fmt.Errorf("internal error: failed to decode json: %v", err)
	}
	fmt.Printf("#%-5d %9.9s %.55s \n\t%s\n", Issue.Number, Issue.User.Login, Issue.Title, Issue.Body)
	return nil
}
