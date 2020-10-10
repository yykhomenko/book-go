package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
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

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("query error: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func GetIssue(owner, repo, number string) (*Issue, error) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get error: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func CreateIssue(owner, repo, title string) (*Issue, error) {
	fields := map[string]string{"title": title}
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(fields); err != nil {
		return nil, err
	}

	client := http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unable to create issue: %s\n", resp.Status)
	}

	var issue Issue
	l := resp.Header.Get("Location")
	i := strings.LastIndex(l, "/")
	n, err := strconv.Atoi(l[i+1:])
	if err != nil {
		return nil, err
	}
	issue.Number = n

	return &issue, nil
}

func UpdateIssue(owner, repo, number string, fields map[string]string) error {
	return patchIssue(owner, repo, number, fields)
}

func CloseIssue(owner, repo, number string) error {
	return patchIssue(owner, repo, number, map[string]string{"state": "close"})
}

func patchIssue(owner, repo, number string, fields map[string]string) error {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(fields); err != nil {
		return err
	}

	client := http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	req, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to update issue: %s\n", resp.Status)
	}

	return nil
}
