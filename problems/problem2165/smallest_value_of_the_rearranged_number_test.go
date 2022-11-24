package problem2165

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int64
	Expected int64
}

var Results = []Result{
	{310, 103},
	{-7605, -7650},
}

func TestRearrangedMinNum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := smallestNumber(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
