// go build -o issues && ./issues search golang
// go build -o issues && ./issues edit OWNER REPO NUMBER
package main

import (
	"fmt"
	"os"
)

var usage = "usage: issues search|create|read|update|close TERM|(OWNER REPO NUMBER)]"

func main() {
	if len(os.Args) < 3 {
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		searchIssues(args)
	} else {
		if len(os.Args) != 5 {
			usageDie()
		}
		owner, repo, number := args[0], args[1], args[2]
		switch cmd {
		case "create":
			createIssue(owner, repo, number)
		case "get":
			getIssue(owner, repo, number)
		case "update":
			updateIssue(owner, repo, number)
		case "close":
			closeIssue(owner, repo, number)
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %q, use: %s\n", cmd, usage)
		}
	}
}

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}
