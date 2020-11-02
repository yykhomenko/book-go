package sorting

import (
	"fmt"
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

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
