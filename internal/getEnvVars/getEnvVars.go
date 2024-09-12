package getenvvars

import "os"

var GitHubToken string

func init() {
	GitHubToken = os.Getenv("GITHUB_TOKEN")
}
