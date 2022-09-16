package problem1710

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Size     int
	Expected int
}

var Results = []Result{
	{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
	{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
}

func TestMaximumUnits(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximumUnits(res.Input, res.Size)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
