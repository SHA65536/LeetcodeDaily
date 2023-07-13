package problem0207

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	Preq     [][]int
	Expected bool
}

var TestCases = []TestCase{
	{2, [][]int{{1, 0}}, true},
	{2, [][]int{{1, 0}, {0, 1}}, false},
	{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
}

func TestCourseSchedule(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := canFinish(tc.N, tc.Preq)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkCourseSchedule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			canFinish(tc.N, tc.Preq)
		}
	}
}
