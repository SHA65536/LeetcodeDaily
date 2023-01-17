package problem0926

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
	{"00110", 1},
	{"010110", 2},
	{"00011000", 2},
}

func TestMonotoneString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minFlipsMonoIncr(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
