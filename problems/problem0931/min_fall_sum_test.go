package problem0931

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
	{[][]int{{-19, 57}, {-40, -5}}, -59},
}

func TestMinFallingPathSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minFallingPathSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
