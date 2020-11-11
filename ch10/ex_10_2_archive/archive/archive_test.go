package archive_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive/archive"
	_ "github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive/archive/tar"
	_ "github.com/yykhomenko/book-gopl/ch10/ex_10_2_archive/archive/zip"
)

func Test_sniff(t *testing.T) {
	f, err := os.Open("tar/testdata/test.tar")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, kind, err := archive.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "tar", kind)
}

func Test_ReadZip(t *testing.T) {
	f, err := os.Open("zip/testdata/test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r, _, err := archive.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}

	w := &bytes.Buffer{}
	if _, err := io.Copy(w, r); err != nil {

	}

	fmt.Println(w.String())
}
