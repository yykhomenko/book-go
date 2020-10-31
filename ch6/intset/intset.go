package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j)) != 0 {
					if buf.Len() > len("}") {
						buf.WriteByte(' ')
					}
					fmt.Fprintf(&buf, "%d", 64*i+j)
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(vs ...int) {
	for _, v := range vs {
		s.Add(v)
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Copy() *IntSet {
	ws := make([]uint64, len(s.words), len(s.words))
	copy(ws, s.words)
	return &IntSet{ws}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Len() (sum int) {
	for _, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j++ {
				if word&(1<<j) != 0 {
					sum++
				}
			}
		}
	}
	return
}

func (s *IntSet) Elems() (elems []uint64) {
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j)) != 0 {
					elems = append(elems, uint64(64*i+j))
				}
			}
		}
	}
	return
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			return
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
