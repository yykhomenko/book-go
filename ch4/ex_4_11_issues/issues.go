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

func readIssue(owner string, repo string, number string) {

}

func updateIssue(owner string, repo string, number string) {

}

func closeIssue(owner string, repo string, number string) {

}
