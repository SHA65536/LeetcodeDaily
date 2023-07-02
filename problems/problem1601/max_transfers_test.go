package problem1601

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	Requests [][]int
	Expected int
}

var TestCases = []TestCase{
	{5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}, 5},
	{3, [][]int{{0, 0}, {1, 2}, {2, 1}}, 3},
	{4, [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}, 4},
	{3, [][]int{{1, 2}, {1, 2}, {2, 2}, {0, 2}, {2, 1}, {1, 1}, {1, 2}}, 4},
}

func TestMaxTransfers(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maximumRequests(tc.N, tc.Requests)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxTransfers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maximumRequests(tc.N, tc.Requests)
		}
	}
}
