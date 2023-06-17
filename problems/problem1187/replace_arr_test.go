package problem1187

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Arr1, Arr2 []int
	Expected   int
}

var TestCases = []TestCase{
	{[]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}, 1},
	{[]int{1, 5, 3, 6, 7}, []int{4, 3, 1}, 2},
	{[]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}, -1},
}

func TestMakeArrayIncreasing(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := makeArrayIncreasing(tc.Arr1, tc.Arr2)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMakeArrayIncreasing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			makeArrayIncreasing(tc.Arr1, tc.Arr2)
		}
	}
}
