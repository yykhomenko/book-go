// go build -o issues && ./issues search golang
// go build -o issues && ./issues edit OWNER REPO NUMBER
// go build -o issues && export GITHUB_USER=USER && export GITHUB_PASS=PASS && ./issues create yykhomenko test
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
		switch cmd {
		case "create":
			createIssue(args[0], args[1], args[2])
		case "get":
			getIssue(args[0], args[1], args[2])
		case "update":
			updateIssue(args[0], args[1], args[2])
		case "close":
			closeIssue(args[0], args[1], args[2])
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %q, use: %s\n", cmd, usage)
		}
	}
}

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}
