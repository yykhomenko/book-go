package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yykhomenko/book-gopl/ch4/github"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: gitserver keyword1 [..., keywordN]")
	}

	searchResult, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	commits, err := github.GetCommits(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	_ = commits
	// users, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if err := indexPage.Execute(w, nil); err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/issues", func(w http.ResponseWriter, _ *http.Request) {
		if err := issuesPage.Execute(w, searchResult); err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/commits", func(w http.ResponseWriter, _ *http.Request) {
		if err := commitsPage.Execute(w, commits); err != nil {
			log.Println(err)
		}
	})

	// http.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
	// 	if err := usersPage.Execute(w, usersResult); err != nil {
	// 		log.Println(err)
	// 	}
	// })

	log.Fatal(http.ListenAndServe(":8000", nil))
}
