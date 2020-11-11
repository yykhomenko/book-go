package archive

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

var ErrFormat = errors.New("archive: unknown format")

// A format holds an image format's name, magic header and how to decode it.
type format struct {
	name, magic string
	magicOffset int
	reader      ArchReader
}

type ArchReader func(f *os.File) (io.Reader, error)

var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

// RegisterFormat registers an archive format for use by NewReader.
// Name is the name of the format, like "tar" or "zip".
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// NewReader is the function that decodes the encoded archive.
// DecodeConfig is the function that decodes just its configuration.
func RegisterFormat(name, magic string, magicOffset int, decode func(*os.File) (io.Reader, error)) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, magic, magicOffset, decode}))
	formatsMu.Unlock()
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(file *os.File) format {
	defer file.Seek(0, io.SeekStart)
	formats, _ := atomicFormats.Load().([]format)
	r := bufio.NewReader(file)
	for _, f := range formats {
		b, err := r.Peek(f.magicOffset + len(f.magic))
		if err == nil && match(f.magic, b[f.magicOffset:]) {
			return f
		}
	}

	return format{}
}

func NewReader(file *os.File) (io.Reader, string, error) {
	format := sniff(file)
	if format.reader == nil {
		return nil, "", ErrFormat
	}
	r, err := format.reader(file)
	return r, format.name, err
}
