package problem2328

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input    [][]int
	Expected int
}

var TestCases = []TestCase{
	{[][]int{{1, 1}, {3, 4}}, 8},
	{[][]int{{1}, {2}}, 3},
	{[][]int{
		{13, 5, 46, 65, 30, 30, 22, 13, 35, 98},
		{46, 80, 8, 17, 26, 29, 73, 91, 49, 70},
		{79, 88, 24, 9, 54, 32, 19, 66, 4, 21},
		{74, 38, 79, 98, 38, 80, 50, 48, 29, 19},
		{72, 43, 49, 42, 12, 99, 94, 75, 5, 96},
		{32, 7, 42, 6, 98, 19, 100, 7, 92, 27},
		{32, 50, 25, 26, 44, 54, 90, 93, 99, 59},
		{96, 65, 55, 49, 16, 12, 66, 16, 80, 5},
		{31, 92, 74, 9, 82, 78, 21, 46, 47, 54},
		{33, 55, 93, 77, 45, 31, 20, 70, 38, 93}}, 716},
}

func TestNumberOfIncreasingPaths(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := countPaths(tc.input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkNumberOfIncreasingPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			countPaths(tc.input)
		}
	}
}
