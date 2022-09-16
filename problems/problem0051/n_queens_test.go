package problem0051

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected [][]string
}

var Results = []Result{
	{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
}

func TestSolveNQueens(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := solveNQueens(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
