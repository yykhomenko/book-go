package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/yykhomenko/book-gopl/ch7/ex_7_8_columns_sort"
)

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, sorting.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, sorting.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, sorting.Length("4m24s")},
}

var tmpl = template.Must(template.New("Tracks").Parse(`
<html><head><title>Tracks</title></head>
<body>
<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
	<th><a href="/tracks?sort=Title">Title</a></th>
	<th><a href="/tracks?sort=Artist">Artist</a></th>
	<th><a href="/tracks?sort=Album">Album</a></th>
	<th><a href="/tracks?sort=Year">Year</a></th>
	<th><a href="/tracks?sort=Length">Length</a></th>
</tr>
{{range .}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

func main() {
	bc := sorting.NewByColumns(tracks, 3)

	http.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		field := strings.ToLower(r.FormValue("sort"))
		switch field {
		case "title":
			bc.Select(sorting.LessTitle)
		case "artist":
			bc.Select(sorting.LessArtist)
		case "album":
			bc.Select(sorting.LessAlbum)
		case "year":
			bc.Select(sorting.LessYear)
		case "length":
			bc.Select(sorting.LessLength)
		}
		sort.Sort(bc)
		tmpl.Execute(w, tracks)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
