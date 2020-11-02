package sorting

type Comparison int
type columnCmp func(a, b *Track) Comparison

const (
	lt Comparison = iota
	eq
	gt
)

func LessTitle(a, b *Track) Comparison {
	switch {
	case a.Title == b.Title:
		return eq
	case a.Title < b.Title:
		return lt
	default:
		return gt
	}
}

func LessArtist(a, b *Track) Comparison {
	switch {
	case a.Artist == b.Artist:
		return eq
	case a.Artist < b.Artist:
		return lt
	default:
		return gt
	}
}

func LessAlbum(a, b *Track) Comparison {
	switch {
	case a.Artist == b.Artist:
		return eq
	case a.Artist < b.Artist:
		return lt
	default:
		return gt
	}
}

func LessYear(a, b *Track) Comparison {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year < b.Year:
		return lt
	default:
		return gt
	}
}

func LessLength(a, b *Track) Comparison {
	switch {
	case a.Length == b.Length:
		return eq
	case a.Length < b.Length:
		return lt
	default:
		return gt
	}
}
