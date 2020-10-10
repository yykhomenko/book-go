package main

import (
	"fmt"
	"log"

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

	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s\n",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func updateIssue(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	// issue to file
	// edit file via editor
	// file to issue

	fields := map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
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
