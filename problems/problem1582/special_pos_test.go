package problem1582

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	Expected int
}

var TestCases = []TestCase{
	{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	{[][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}, 1},
}

func TestSpecialPositionMatrix(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var got = numSpecial(tc.Input)
		var want = tc.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSpecialPositionMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numSpecial(tc.Input)
		}
	}
}
