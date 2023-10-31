package problem2433

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Expected []int
}

var TestCases = []TestCase{
	{[]int{5, 2, 0, 3, 1}, []int{5, 7, 2, 3, 2}},
	{[]int{13}, []int{13}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 3, 1, 7, 1, 3, 1, 15, 1, 3}},
}

func TestOriginalXor(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		var cpy = make([]int, len(tc.Input))
		copy(cpy, tc.Input)
		got := findArray(cpy)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkOriginalXor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			var cpy = make([]int, len(tc.Input))
			copy(cpy, tc.Input)
			findArray(cpy)
		}
	}
}
