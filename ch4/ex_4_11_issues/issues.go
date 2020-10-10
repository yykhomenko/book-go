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

func createIssue(owner string, repo string, number string) {

}

func getIssue(owner string, repo string, number string) {
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

func updateIssue(owner string, repo string, number string) {

}

func closeIssue(owner string, repo string, number string) {

}
