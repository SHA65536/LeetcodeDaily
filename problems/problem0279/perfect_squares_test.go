package problem0279

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
	{12, 3},
	{13, 2},
	{48, 3},
	{9999, 4},
}

func TestPerfectSquares(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numSquares(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
