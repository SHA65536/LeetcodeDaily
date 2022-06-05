package problem0052

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
	{1, 1},
	{2, 0},
	{3, 0},
	{4, 2},
	{5, 10},
	{6, 4},
	{7, 40},
	{8, 92},
	{9, 352},
}

func TestTotalNQueens(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := totalNQueens(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
