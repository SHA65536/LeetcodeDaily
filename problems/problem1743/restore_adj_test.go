package problem1743

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	Expected []int
}

var TestCases = []TestCase{
	{[][]int{{2, 1}, {3, 4}, {3, 2}}, []int{1, 2, 3, 4}},
	{[][]int{{4, -2}, {1, 4}, {-3, 1}}, []int{-2, 4, 1, -3}},
	{[][]int{{100000, -100000}}, []int{100000, -100000}},
}

func TestRestoreAdjacent(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		got := restoreArray(tc.Input)
		assert.True(elementsMatchOneOf(got, want, reverse(want)), fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRestoreAdjacent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			restoreArray(tc.Input)
		}
	}
}

func reverse(in []int) []int {
	var res = make([]int, len(in))

	for i := range in {
		res[len(in)-i-1] = in[i]
	}

	return res
}

func elementsMatchOneOf(arr []int, options ...[]int) bool {
	for _, option := range options {
		if assert.ElementsMatch(nil, arr, option) {
			return true
		}
	}
	return false
}
