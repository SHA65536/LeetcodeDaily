package problem1095

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Target   int
	Expected int
}

var TestCases = []TestCase{
	{[]int{1, 2, 3, 4, 5, 3, 1}, 3, 2},
	{[]int{0, 1, 2, 4, 2, 1}, 3, -1},
}

func TestFindInMountain(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		var m = MountainArray(tc.Input)
		got := findInMountainArray(tc.Target, &m)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFindInMountain(b *testing.B) {
	for _, tc := range TestCases {
		for i := 0; i < b.N; i++ {
			var m = MountainArray(tc.Input)
			findInMountainArray(tc.Target, &m)
		}
	}
}
