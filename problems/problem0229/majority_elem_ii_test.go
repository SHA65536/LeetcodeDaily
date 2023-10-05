package problem0229

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
	{[]int{3, 2, 3}, []int{3}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
}

func TestMajorityElementII(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := majorityElement(tc.Input)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMajorityElementII(b *testing.B) {
	for _, tc := range TestCases {
		for i := 0; i < b.N; i++ {
			majorityElement(tc.Input)
		}
	}
}
