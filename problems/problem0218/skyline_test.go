package problem0218

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected [][]int
}

var Results = []Result{
	{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}, [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
	{[][]int{{0, 2, 3}, {2, 5, 3}}, [][]int{{0, 3}, {5, 0}}},
}

func TestSkylineProblem(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := getSkyline(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
