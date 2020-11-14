package intset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.Equal(t, "{2 4}", x.String())
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.AddAll(3, 4)
	assert.True(t, x.Has(3))
	assert.True(t, x.Has(4))
}

func TestHas(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.True(t, x.Has(4))
	assert.True(t, !x.Has(3))
}

func TestRemove(t *testing.T) {
	x := IntSet{}
	v := 4
	x.Add(v)
	x.Remove(v)
	assert.False(t, x.Has(v))
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	y := x.Copy()
	assert.True(t, y.Has(1))
	assert.True(t, y.Has(2))
}

func TestClear(t *testing.T) {
	x := IntSet{}
	x.Add(1)
	x.Add(3)
	x.Add(5)
	x.Clear()
	assert.Equal(t, 0, x.Len())
}

func TestLen(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.Equal(t, 2, x.Len())
}

func TestElems(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.Equal(t, []uint{2, 4}, x.Elems())
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.UnionWith(&y)
	assert.Equal(t, "{1 9 42 144}", x.String())
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(32)
	x.Add(200)
	x.Add(300)
	y.Add(300)
	y.Add(400)
	x.IntersectWith(&y)
	assert.Equal(t, "{300}", x.String())
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	y.Add(1)
	y.Add(4)
	y.Add(5)
	z := x.Copy()
	x.DifferenceWith(&y)
	assert.Equal(t, "{2 3}", x.String())
	y.DifferenceWith(z)
	assert.Equal(t, "{4 5}", y.String())
}

func TestSymmetricDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	y.Add(1)
	y.Add(4)
	y.Add(5)
	x.SymmetricDifferenceWith(&y)
	assert.Equal(t, "{2 3 4 5}", x.String())
}
