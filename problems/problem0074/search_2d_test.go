package problem0074

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	Target   int
	Expected bool
}

var TestCases = []TestCase{
	{
		Input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		Target:   3,
		Expected: true,
	},
	{
		Input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		Target:   13,
		Expected: false,
	},
}

func TestSearchIn2DMatrix(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := searchMatrix(tc.Input, tc.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSearchIn2DMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			searchMatrix(tc.Input, tc.Target)
		}
	}
}
