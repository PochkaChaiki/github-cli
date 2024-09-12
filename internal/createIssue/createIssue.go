package createIssue

import (
	"bytes"
	"encoding/json"
	"fmt"
	getenvvars "githubIssues/internal/getEnvVars"
	"log"
	"net/http"
	"strings"
)

const githubAPI = "https://api.github.com/repos"

func CreateIssue(owner string, repo string, title string, text string) error {

	body := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{
		Title: title,
		Body:  text,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return fmt.Errorf("internal error in json encoder: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, strings.Join([]string{githubAPI, owner, repo, "issues"}, "/"), &buf)
	if err != nil {
		return fmt.Errorf("internal error: failed to create request: %v", err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+getenvvars.GitHubToken)

	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("internal error: failed to send request: %v", err)
	}

	log.Printf("Status: %d", resp.StatusCode)
	return nil
}
