package githubissues

import (
	"bytes"
	"encoding/json"
	"fmt"
	getenvvars "github-cli/internal/getEnvVars"
	"net/http"
	"strings"
)

func CreateIssue(owner string,
	repo string,
	title string,
	text string,
	assignee string,
	milestone string,
	assignees []string) (int, error) {

	body := Issue{
		Title:     title,
		Body:      text,
		Assignee:  assignee,
		Assignees: assignees,
		Milestone: milestone,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return 0, fmt.Errorf("internal error in json encoder: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, strings.Join([]string{getenvvars.GitHubAPI, owner, repo, "issues"}, "/"), &buf)
	if err != nil {
		return 0, fmt.Errorf("internal error: failed to create request: %v", err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+getenvvars.GitHubToken)

	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return 0, fmt.Errorf("internal error: failed to send request: %v", err)
	}

	// log.Printf("Status: %d", resp.StatusCode)
	return resp.StatusCode, nil
}
