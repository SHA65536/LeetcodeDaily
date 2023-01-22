package problem0131

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected [][]string
}

var Results = []Result{
	{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	{"a", [][]string{{"a"}}},
	{"aaa", [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}}},
	{"brjwevkhxozncuzy", [][]string{{"b", "r", "j", "w", "e", "v", "k", "h", "x", "o", "z", "n", "c", "u", "z", "y"}}},
	{"kpethkuaolsjsody", [][]string{{"k", "p", "e", "t", "h", "k", "u", "a", "o", "l", "s", "j", "s", "o", "d", "y"}, {"k", "p", "e", "t", "h", "k", "u", "a", "o", "l", "sjs", "o", "d", "y"}}},
	{"jjacqmughpbxpfhp", [][]string{{"j", "j", "a", "c", "q", "m", "u", "g", "h", "p", "b", "x", "p", "f", "h", "p"}, {"jj", "a", "c", "q", "m", "u", "g", "h", "p", "b", "x", "p", "f", "h", "p"}}},
	{"sxuhxdlqybxamsmz", [][]string{{"s", "x", "u", "h", "x", "d", "l", "q", "y", "b", "x", "a", "m", "s", "m", "z"}, {"s", "x", "u", "h", "x", "d", "l", "q", "y", "b", "x", "a", "msm", "z"}}},
	{"aaaaa", [][]string{{"a", "a", "a", "a", "a"}, {"a", "a", "a", "aa"}, {"a", "a", "aa", "a"}, {"a", "a", "aaa"}, {"a", "aa", "a", "a"}, {"a", "aa", "aa"}, {"a", "aaa", "a"}, {"a", "aaaa"}, {"aa", "a", "a", "a"}, {"aa", "a", "aa"}, {"aa", "aa", "a"}, {"aa", "aaa"}, {"aaa", "a", "a"}, {"aaa", "aa"}, {"aaaa", "a"}, {"aaaaa"}}},
}

func TestPartitionPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partition(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
func BenchmarkPartitionPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			partition(res.Input)
		}
	}
}

func TestPartitionPalindromeCache(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partitionCache(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
func BenchmarkPartitionPalindromeCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			partitionCache(res.Input)
		}
	}
}
