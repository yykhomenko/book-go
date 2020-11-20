package intset

import (
	"bytes"
	"fmt"
)

// ex 6.5 N is number of uint bits on current platform
const N = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < N; j++ {
				if word&(1<<uint(j)) != 0 {
					if buf.Len() > len("}") {
						buf.WriteByte(' ')
					}
					fmt.Fprintf(&buf, "%d", N*i+j)
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/N, uint(x%N)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/N, uint(x%N)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// ex 6.2
func (s *IntSet) AddAll(vs ...int) {
	for _, v := range vs {
		s.Add(v)
	}
}

// ex 6.1
func (s *IntSet) Remove(x int) {
	word, bit := x/N, uint(x%N)
	s.words[word] &^= 1 << bit
}

// ex 6.1
func (s *IntSet) Copy() *IntSet {
	ws := make([]uint, len(s.words), len(s.words))
	copy(ws, s.words)
	return &IntSet{ws}
}

// ex 6.1
func (s *IntSet) Clear() {
	s.words = nil
}

// ex 6.1
func (s *IntSet) Len() (sum int) {
	for _, word := range s.words {
		if word != 0 {
			for j := 0; j < N; j++ {
				if word&(1<<j) != 0 {
					sum++
				}
			}
		}
	}
	return
}

// ex 6.4
func (s *IntSet) Elems() (elems []uint) {
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < N; j++ {
				if word&(1<<uint(j)) != 0 {
					elems = append(elems, uint(N*i+j))
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

// ex 6.3
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			return
		}
	}
}

// ex 6.3
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// ex 6.3
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
