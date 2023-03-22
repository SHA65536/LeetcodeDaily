package problem2492

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	N        int
	Expected int
}

var Results = []Result{
	{[][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}, 4, 5},
	{[][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}, 4, 2},
}

func TestShortestDistanceBetweenCities(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minScore(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
