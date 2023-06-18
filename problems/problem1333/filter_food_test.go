package problem1333

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	v, p, d  int
	Expected []int
}

var TestCases = []TestCase{
	{[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 1, 50, 10, []int{3, 1, 5}},
	{[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 50, 10, []int{4, 3, 2, 1, 5}},
	{[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 30, 3, []int{4, 5}},
}

func TestFilterRestaurants(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := filterRestaurants(tc.Input, tc.v, tc.p, tc.d)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFilterRestaurants(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			filterRestaurants(tc.Input, tc.v, tc.p, tc.d)
		}
	}
}
