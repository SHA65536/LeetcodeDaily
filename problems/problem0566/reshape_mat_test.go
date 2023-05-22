package problem0566

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	R, C     int
	Expected [][]int
}

var Results = []Result{
	{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
	{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
	{[][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 4, 4, [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}},
	{[][]int{{1, 2}, {3, 4}}, 2, 3, [][]int{{1, 2}, {3, 4}}},
}

func TestReshapeMat(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := matrixReshape(res.Input, res.R, res.C)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkReshapeMat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			matrixReshape(res.Input, res.R, res.C)
		}
	}
}
