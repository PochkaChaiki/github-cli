package githubissues

import (
	"fmt"
	"testing"
)

var TestNumber = 1

func TestCreateIssue(t *testing.T) {
	_, err := CreateIssue("PochkaChaiki", "github-cli", "TestIssue"+fmt.Sprint(TestNumber), "TestIssue", "PochkaChaiki", "", nil)
	if err != nil {
		t.Errorf("CreateIssue() = %v", err)
	}
}

func TestGetListOfIssues(t *testing.T) {
	_, err := GetListOfIssues("PochkaChaiki", "github-cli")
	if err != nil {
		t.Errorf("GetListOfIssues() = %v", err)
	}
}

func TestGetIssue(t *testing.T) {
	_, err := GetIssue("PochkaChaiki", "github-cli", 3)
	if err != nil {
		t.Errorf("GetIssue() = %v", err)
	}
}

func TestUpdateIssue(t *testing.T) {
	_, err := UpdateIssue("PochkaChaiki", "github-cli", 3, "NewTestIssue"+fmt.Sprint(TestNumber), "NewTestIssue", "PochkaChaiki", nil, "", "", "")
	if err != nil {
		t.Errorf("UpdateIssue() = %v", err)
	}
}
