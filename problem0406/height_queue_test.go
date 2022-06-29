package problem0406

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
	{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
		[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
	{[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}},
		[][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}},
	{[][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}},
		[][]int{{3, 0}, {6, 0}, {7, 0}, {5, 2}, {3, 4}, {5, 3}, {6, 2}, {2, 7}, {9, 0}, {1, 9}}},
}

func TestReconstructQueue(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		copied := make([][]int, len(res.Input))
		copy(copied, res.Input)
		want := res.Expected
		got := reconstructQueue(copied)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
