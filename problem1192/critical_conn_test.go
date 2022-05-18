package problem1192

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	N        int
	Expected [][]int
}

var Results = []Result{
	{
		Input:    [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}},
		N:        4,
		Expected: [][]int{{1, 3}},
	},
	{
		Input:    [][]int{{0, 1}},
		N:        2,
		Expected: [][]int{{0, 1}},
	},
}

func TestCriticalConnections(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := criticalConnections(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
