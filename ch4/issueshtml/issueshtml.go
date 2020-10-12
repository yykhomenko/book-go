// go run issueshtml.go github > out.html
package main

import (
	"html/template"
	"log"
	"os"

	"github.com/yykhomenko/book-gopl/ch4/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} themes</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: issueshtml keyword1 [..., keywordN] > out.html")
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
