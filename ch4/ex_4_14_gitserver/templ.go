package main

import "html/template"

var indexPage = template.Must(template.New("index").Parse(`
<h1>Gitserver</h1>
<table>
<tr><td><a href='/issues'>Issues</a></td></tr>
<tr><td><a href='/commits'>Commits</a></td></tr>
<tr><td><a href='/users'>Users</a></td></tr>
</table>
`))

var issuesPage = template.Must(template.New("issues").Parse(`
<h1>{{.TotalCount}} themes</h1>
<a href='/'>Index</a>
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

var usersPage = template.Must(template.New("users").Parse(`
<h1>{{.TotalCount}} themes</h1>
<a href='/'>Index</a>
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
