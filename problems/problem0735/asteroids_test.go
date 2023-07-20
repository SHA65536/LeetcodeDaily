package problem0735

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
	{[]int{5, 10, -5}, []int{5, 10}},
	{[]int{8, -8}, []int{}},
	{[]int{10, 2, -5}, []int{10}},
}

func TestAsteroidCollisions(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := asteroidCollision(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkAsteroidCollisions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			asteroidCollision(tc.Input)
		}
	}
}
