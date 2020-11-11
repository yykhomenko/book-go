package tar

import (
	"io"

	archive "github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive"
)

func Reader(r io.Reader) (io.Reader, error) {
	return r, nil
}

func init() {
	archive.RegisterFormat("zip", "PK", 0, Reader)
}
