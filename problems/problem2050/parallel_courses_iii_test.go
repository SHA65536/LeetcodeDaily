package problem2050

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N         int
	Relations [][]int
	Time      []int
	Expected  int
}

var TestCases = []TestCase{
	{3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}, 8},
	{5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5}, 12},
}

func TestParallelCoursesIII(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minimumTime(tc.N, tc.Relations, tc.Time)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkVParallelCoursesIII(b *testing.B) {
	for _, tc := range TestCases {
		for i := 0; i < b.N; i++ {
			minimumTime(tc.N, tc.Relations, tc.Time)
		}
	}
}
