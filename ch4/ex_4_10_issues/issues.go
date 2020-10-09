// go run issues.go github
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yykhomenko/book-gopl/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d themes\n", result.TotalCount)

	fmt.Println("up to a month:")
	printBetween(result, time.Now().AddDate(0, -1, 0), time.Now())
	fmt.Println("less then a year:")
	printBetween(result, time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, -1, 0))
	fmt.Println("over a year:")
	printBetween(result, time.Time{}, time.Now().AddDate(-1, 0, 0))
}

func printBetween(issues *github.IssuesSearchResult, a, b time.Time) {
	for _, item := range issues.Items {
		if a.Before(item.CreatedAt) && b.After(item.CreatedAt) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}
