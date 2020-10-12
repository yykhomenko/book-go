package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const CommitsURLFormat = "https://api.github.com/repos/%s/%s/commits"

type CommitInfo struct {
	SHA     string
	Commit  *Commit
	Author  *User
	HTMLURL string `json:"html_url"`
}

type Committer struct {
	Author      *User
	CommitCount int
}

type Commit struct {
	Author  Person
	Message string
}

type Person struct {
	Date time.Time
}

func GetCommits(owner, repo string) ([]CommitInfo, error) {
	url := fmt.Sprintf(CommitsURLFormat, owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get error: %s", resp.Status)
	}

	var commits []CommitInfo
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, err
	}

	return commits, nil
}
