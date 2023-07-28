package problem0486

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        []int
	Expected bool
}

var TestCases = []TestCase{
	{[]int{1, 5, 2}, false},
	{[]int{1, 5, 233, 7}, true},
}

func TestPredictTheWinner(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := PredictTheWinner(tc.N)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkPredictTheWinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			PredictTheWinner(tc.N)
		}
	}
}
