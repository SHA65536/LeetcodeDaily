package problem0028

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Target   string
	Expected int
}

var Results = []Result{
	{"hello", "ll", 2},
	{"aaaaa", "bba", -1},
	{"hello", "", 0},
}

func TestStrStr(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := strStr(res.Input, res.Target)
		want := res.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
