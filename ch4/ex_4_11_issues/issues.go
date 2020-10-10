package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/yykhomenko/book-gopl/ch4/github"
)

func searchIssues(args []string) {
	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d themes\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func createIssue(owner, repo, title string) {
	issue, err := github.CreateIssue(owner, repo, title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created #%d\n", issue.Number)
}

func getIssue(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}

	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\ncomment: %s\n",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func getIssues(owner, repo string) {
	issues, err := github.GetIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range issues {
		body := issue.Body
		if body == "" {
			body = "<empty>\n"
		}

		fmt.Printf("repo: %s/%s\nnumber: %d\nuser: %s\ntitle: %s\ncomment: %s\n",
			owner, repo, issue.Number, issue.User.Login, issue.Title, body)
	}
}

func updateIssue(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	}

	content, err := json.MarshalIndent(fields, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	tmpfile, err := ioutil.TempFile("", "issues-*")
	if err != nil {
		log.Fatal(err)
	}

	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	tmpfile.Sync()

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tmpfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	tmpfile.Seek(0, 0)
	content, err = ioutil.ReadAll(tmpfile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(content, &fields); err != nil {
		log.Fatal(err)
	}

	if err := github.UpdateIssue(owner, repo, number, fields); err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated")
}

func closeIssue(owner, repo, number string) {
	if err := github.CloseIssue(owner, repo, number); err != nil {
		log.Fatal(err)
	}
	fmt.Println("closed")
}
