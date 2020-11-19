package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v\n", test.input, got)
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("random generator seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("random generator seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		fmt.Println(p)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

func randomPalindrome(rnd *rand.Rand) string {
	n := rnd.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rnd.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rnd *rand.Rand) string {
	n := rnd.Intn(23) + 2
	runes := make([]rune, n)
	for i := 0; i < n; {
		r := rune(rnd.Intn(127-32) + 32)

		if unicode.IsLetter(r) ||
			unicode.IsPunct(r) ||
			unicode.IsSpace(r) {
			runes[i] = r
			i++
		}
	}

	letters := onlyLetters(runes)

	if len(letters) == 0 ||
		unicode.ToLower(letters[0]) ==
			unicode.ToLower(letters[len(letters)-1]) {
		return randomNonPalindrome(rnd)
	}

	return string(runes)
}

func onlyLetters(in []rune) (out []rune) {
	for _, r := range in {
		if unicode.IsLetter(r) {
			out = append(out, r)
		}
	}
	return
}
