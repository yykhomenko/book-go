package archive_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	archive "github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive"
	_ "github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive/tar"
	_ "github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive/zip"
)

func Test_sniff(t *testing.T) {
	f, err := os.Open("tar/testdata/test.tar")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, kind, err := archive.Reader(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "tar", kind)
}
