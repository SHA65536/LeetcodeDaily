package problem0547

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
	{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
	{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
}

func TestNumberOfProvinces(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := findCircleNum(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkNumberOfProvinces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findCircleNum(tc.Input)
		}
	}
}
