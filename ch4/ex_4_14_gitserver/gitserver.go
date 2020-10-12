// go build -o gitserver && ./gitserver yykhomenko test
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

	searchResult, err := github.GetIssues(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal("unable get issues:", err)
	}

	commits, err := github.GetCommits(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal("unable get commits:", err)
	}

	// authorFreq := make(map[github.User]int)
	// for _, commit := range commits {
	// 	authorFreq[commit.Author]++
	// }
	//
	// var committers []github.Committer
	// for login, count := range authorFreq {
	// 	committer := github.Committer{
	// 		Author:      nil,
	// 		CommitCount: 0,
	// 	}
	// 	committers = append(committers, )
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

	http.HandleFunc("/committers", func(w http.ResponseWriter, _ *http.Request) {
		if err := committersPage.Execute(w, committersPage); err != nil {
			log.Println(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
