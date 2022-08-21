package problem0936

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Stamp, Target string
	Expected      []int
}

var Results = []Result{
	{"abc", "ababc", []int{0, 2}},
	{"abca", "aabcaca", []int{3, 0, 1}},
}

func TestStampString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := movesToStamp(res.Stamp, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
