package str

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(sort.StringSlice([]string{"g", "o", "o", "g"})))
	assert.True(t, IsPalindrome(sort.IntSlice([]int{1, 2, 3, 2, 1})))
	assert.False(t, IsPalindrome(sort.IntSlice([]int{1, 2, 3, 4, 5})))
}
