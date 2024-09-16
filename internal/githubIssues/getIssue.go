package githubissues

import (
	"encoding/json"
	"fmt"
	getenvvars "github-cli/internal/getEnvVars"
	"net/http"
	"strings"
)

func GetListOfIssues(owner string, repo string) ([]IssueReturn, error) {

	req, err := http.NewRequest(http.MethodGet, strings.Join([]string{getenvvars.GitHubAPI, owner, repo, "issues"}, "/"), nil)
	if err != nil {
		return nil, fmt.Errorf("internal error: failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Beaver "+getenvvars.GitHubToken)
	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("internal error: failed to send request: %v", err)
	}
	var IssuesResult []IssueReturn
	err = json.NewDecoder(resp.Body).Decode(&IssuesResult)
	if err != nil {
		return nil, fmt.Errorf("internal error: failed to decode json: %v", err)
	}

	return IssuesResult, nil
}

func GetIssue(owner string, repo string, number int) (*IssueReturn, error) {

	req, err := http.NewRequest(http.MethodGet, strings.Join([]string{getenvvars.GitHubAPI, owner, repo, "issues", fmt.Sprint(number)}, "/"), nil)
	if err != nil {
		return nil, fmt.Errorf("internal error: failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Beaver "+getenvvars.GitHubToken)
	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("internal error: failed to send request: %v", err)
	}
	var Issue IssueReturn
	err = json.NewDecoder(resp.Body).Decode(&Issue)
	if err != nil {
		return nil, fmt.Errorf("internal error: failed to decode json: %v", err)
	}
	return &Issue, nil
}
