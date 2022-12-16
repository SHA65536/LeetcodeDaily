package problem1137

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{4, 4},
	{25, 1389537},
}

func TestTribonacci(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := tribonacci(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
