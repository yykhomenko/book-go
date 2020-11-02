package sorting

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"text/tabwriter"
)

func (c *ByColumns) PrintTracks() {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range c.tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

var html = `
<html><head><title>Tracks</title></head>
<body>
<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
	<th>Title</th>
	<th>Artist</th>
	<th>Album</th>
	<th>Year</th>
	<th>Length</th>
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
`
var tmpl = template.Must(template.New("Tracks").Parse(html))

func (c *ByColumns) PrintTracksHTML(w io.Writer) {
	tmpl.Execute(w, c.tracks)
}
