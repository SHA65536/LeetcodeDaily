package problem0462

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{1, 2, 3}, 2},
	{[]int{1, 10, 2, 9}, 16},
}

func TestMinMoves2(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minMoves2(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
