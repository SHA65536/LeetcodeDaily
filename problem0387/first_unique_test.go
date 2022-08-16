package problem0387

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"leetcode", 0},
	{"loveleetcode", 2},
	{"aabb", -1},
	{"a", 0},
	{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", 128},
}

func TestFirstUniqChar(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := firstUniqChar(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFirstUniqChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			firstUniqChar(res.Input)
		}
	}
}

func TestFirstUniqCharMap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := firstUniqCharMap(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFirstUniqCharMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			firstUniqCharMap(res.Input)
		}
	}
}
