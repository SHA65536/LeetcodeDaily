package problem1544

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"leEeetcode", "leetcode"},
	{"abBAcC", ""},
	{"s", "s"},
}

func TestMakeStringGood(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := makeGood(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMakeStringGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			makeGood(res.Input)
		}
	}
}
