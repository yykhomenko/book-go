package ex_4_6_onespace

import (
	"unicode"
	"unicode/utf8"
)

func onespace(bs []byte) []byte {
	var lastSpace bool
	var rp, wp int

	for rp < (len(bs)) {
		r, size := utf8.DecodeRune(bs[rp:])
		if !unicode.IsSpace(r) {
			utf8.EncodeRune(bs[wp:], r)
			wp += size
			lastSpace = false
		} else {
			if !lastSpace {
				size = utf8.EncodeRune(bs[wp:], ' ')
				wp += size
			}
			lastSpace = true
		}
		rp += size
	}

	return bs[:wp]
}
