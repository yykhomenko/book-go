package intset_test

import (
	"math/rand"
	"testing"

	"github.com/yykhomenko/book-gopl/ch6/intset"
)

const maxRandInt = 1e8

var IntSet = new(intset.IntSet)
var MapSet = make(map[int]struct{})
var Has bool
var Len int

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSet.Add(rand.Intn(maxRandInt))
	}
}

func BenchmarkAddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSet.AddAll(rand.Intn(maxRandInt))
	}
}

func BenchmarkHas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Has = IntSet.Has(rand.Intn(maxRandInt))
	}
}

func BenchmarkLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len = IntSet.Len()
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSet = IntSet.Copy()
	}
}

func BenchmarkRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSet.Remove(rand.Intn(maxRandInt))
	}
}

func BenchmarkMapAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapSet[rand.Intn(maxRandInt)] = struct{}{}
	}
}

func BenchmarkMapHas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, Has = MapSet[rand.Intn(maxRandInt)]
	}
}

func BenchmarkMapLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len = len(MapSet)
	}
}

func BenchmarkMapRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		delete(MapSet, rand.Intn(maxRandInt))
	}
}
