package problem1569

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Expected int
}

var TestCases = []TestCase{
	{[]int{2, 1, 3}, 1},
	{[]int{3, 4, 5, 1, 2}, 5},
	{[]int{1, 2, 3}, 0},
}

func TestComboBst(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := numOfWays(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkComboBst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numOfWays(tc.Input)
		}
	}
}
