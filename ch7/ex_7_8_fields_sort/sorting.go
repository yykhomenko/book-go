package sorting

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var sorters = map[fieldCmp]func(_, _ *Track) bool{
	{"Artist", true}: func(i, j *Track) bool {
		return i.Artist < j.Artist
	},
	{"Artist", false}: func(i, j *Track) bool {
		return i.Artist > j.Artist
	},
}

type fieldCmp struct {
	name string
	less bool
}

type fieldsSort struct {
	t      []*Track
	fields []fieldCmp
}

func (x fieldsSort) Len() int { return len(x.t) }
func (x fieldsSort) Less(i, j int) bool {
	for i := len(x.fields) - 1; i >= 0; i-- {
		key := x.fields[i]
		res := sorters[key](x.t[i], x.t[j])
		if res {
			return true
		}
	}
	return false
}
func (x fieldsSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x fieldsSort) PrintTracks() {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range x.t {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func (x fieldsSort) ArtistFieldLess(i, j *Track) bool {
	return i.Artist < j.Artist
}

func (x fieldsSort) ArtistFieldGreat(i, j *Track) bool {
	return i.Artist > j.Artist
}

func (x *fieldsSort) AddBy(fieldName string, less bool) {
	key := fieldCmp{fieldName, less}

	if x.fields == nil {
		x.fields = append(x.fields, key)
	} else if len(x.fields) > 0 && x.fields[0] == key {
		x.fields[0].less = !x.fields[0].less
	} else {
		x.fields = append(x.fields, key)
		if len(x.fields) > 3 {
			copy(x.fields, x.fields[len(x.fields)-3:])
			x.fields = x.fields[:3]
		}
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
