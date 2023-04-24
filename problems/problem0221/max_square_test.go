package problem0221

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]byte
	Expected int
}

var Results = []Result{
	{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 4},
	{[][]byte{{'0', '1'}, {'1', '0'}}, 1},
	{[][]byte{{'0'}}, 0},
}

func TestBiggestSquare(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximalSquare(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", got))
	}
}
