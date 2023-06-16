package problem2369

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Expected bool
}

var TestCases = []TestCase{
	{[]int{4, 4, 4, 5, 6}, true},
	{[]int{1, 1, 1, 2}, false},
}

func TestCheckValidPartition(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := validPartition(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkCheckValidPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			validPartition(tc.Input)
		}
	}
}
