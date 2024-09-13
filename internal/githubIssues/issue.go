package githubissues

import "time"

type User struct {
	Login   string `json:"login,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}

type Issue struct {
	Number      int    `json:"number,omitempty"`
	HTMLURL     string `json:"html_url,omitempty"`
	Title       string `json:"title,omitempty"`
	State       string `json:"state,omitempty"`
	StateReason string `json:"state_reason,omitempty"`
	User        *User
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Body        string    `json:"body,omitempty"`
	Assignee    string    `json:"assignee,omitempty"`
	Milestone   string    `json:"milestone,omitempty"`
	Labels      []string  `json:"labels,omitempty"`
	Assignees   []string  `json:"assignees,omitempty"`
}
