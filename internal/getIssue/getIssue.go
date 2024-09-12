package getissue

import (
	"encoding/json"
	"fmt"
	getenvvars "githubIssues/internal/getEnvVars"
	"net/http"
	"strings"
	"time"
)

const githubAPI = "https://api.github.com/repos"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

func GetListOfIssues(owner string, repo string) error {

	req, err := http.NewRequest(http.MethodGet, strings.Join([]string{githubAPI, owner, repo, "issues"}, "/"), nil)
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

	req, err := http.NewRequest(http.MethodGet, strings.Join([]string{githubAPI, owner, repo, "issues", fmt.Sprint(number)}, "/"), nil)
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
