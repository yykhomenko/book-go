package popcount_test

import (
	"testing"

	p24 "github.com/yykhomenko/book-gopl/ch2/ex_2_4_popcount"
	p25 "github.com/yykhomenko/book-gopl/ch2/ex_2_5_popcount"
	p262 "github.com/yykhomenko/book-gopl/ch2/popcount"
)

var Result int

func BenchmarkPopCount262(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = p262.PopCount(123)
	}
}

func BenchmarkPopCount24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = p24.PopCount(123)
	}
}

func BenchmarkPopCount25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = p25.PopCount(123)
	}
}
