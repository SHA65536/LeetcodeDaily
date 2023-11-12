package problem0815

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Routes   [][]int
	Src, Dst int
	Expected int
}

var TestCases = []TestCase{
	{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6, 2},
	{[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12, -1},
}

func TestBusRoutes(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		var got = numBusesToDestination(tc.Routes, tc.Src, tc.Dst)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkBusRoutes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numBusesToDestination(tc.Routes, tc.Src, tc.Dst)
		}
	}
}
