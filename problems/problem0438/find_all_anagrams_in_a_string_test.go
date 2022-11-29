package problem0438

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S, P     string
	Expected []int
}

var Results = []Result{
	{"cbaebabacd", "abc", []int{0, 6}},
	{"abcbebsbdbebugfjbebbcbebcacbebbcbacbcbacbbcbabcbccbababcbacbcbabcabacbacbacbacbabc", "abc", []int{0, 23, 25, 31, 32, 33, 36, 37, 38, 42, 44, 49, 53, 55, 56, 57, 60, 62, 63, 64, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 79}},
	{"abab", "ab", []int{0, 1, 2}},
}

func TestFindAllAnagrams(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findAnagrams(res.S, res.P)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindAllAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findAnagrams(res.S, res.P)
		}
	}
}

func TestFindAllAnagramsHash(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findAnagramsHash(res.S, res.P)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindAllAnagramsHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findAnagramsHash(res.S, res.P)
		}
	}
}
