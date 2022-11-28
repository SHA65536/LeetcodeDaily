package problem2225

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
	{[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
		[][]int{{1, 2, 10}, {4, 5, 7, 8}}},
	{[][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}},
		[][]int{{1, 2, 5, 6}, {}}},
}

func TestFindWinners(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findWinners(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindWinners(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findWinners(res.Input)
		}
	}
}
