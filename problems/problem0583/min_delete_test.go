package problem0583

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S1       string
	S2       string
	Expected int
}

var Results = []Result{
	{"sea", "eat", 2},
	{"leetcode", "etco", 4},
}

func TestMinDistance(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := minDistance(res.S1, res.S2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
