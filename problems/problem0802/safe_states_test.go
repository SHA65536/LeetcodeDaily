package problem0802

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
	{[][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}, []int{2, 4, 5, 6}},
	{[][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}, []int{4}},
}

func TestEventualSafeStates(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := eventualSafeNodes(tc.Input)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkEventualSafeStates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			eventualSafeNodes(tc.Input)
		}
	}
}
