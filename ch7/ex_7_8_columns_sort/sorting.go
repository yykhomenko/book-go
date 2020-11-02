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

func (t Track) String() string {
	return fmt.Sprintf("%s-%s(%d)", t.Title, t.Artist, t.Year)
}

type ByColumns struct {
	tracks     []*Track
	columns    []columnCmp
	maxColumns int
}

func NewByColumns(tracks []*Track, maxColumns int) *ByColumns {
	return &ByColumns{tracks, nil, maxColumns}
}

func (c *ByColumns) Len() int      { return len(c.tracks) }
func (c *ByColumns) Swap(i, j int) { c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i] }

func (c *ByColumns) Less(i, j int) bool {
	for _, f := range c.columns {
		cmp := f(c.tracks[i], c.tracks[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (c *ByColumns) Select(cmp columnCmp) {
	c.columns = append([]columnCmp{cmp}, c.columns...)
	if len(c.columns) > c.maxColumns {
		c.columns = c.columns[:c.maxColumns]
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

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
