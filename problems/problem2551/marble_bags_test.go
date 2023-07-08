package problem2551

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	K        int
	Expected int64
}

var TestCases = []TestCase{
	{[]int{1, 3, 5, 1}, 2, 4},
	{[]int{1, 3}, 2, 0},
}

func TestMarbleBags(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := putMarbles(tc.Input, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMarbleBags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			putMarbles(tc.Input, tc.K)
		}
	}
}
