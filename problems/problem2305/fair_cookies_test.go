package problem2305

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	K        int
	Expected int
}

var TestCases = []TestCase{
	{[]int{8, 15, 10, 20, 8}, 2, 31},
	{[]int{6, 1, 3, 2, 2, 4, 1, 2}, 3, 7},
}

func TestFairCookieDistribution(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := distributeCookies(tc.Input, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFairCookieDistribution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			distributeCookies(tc.Input, tc.K)
		}
	}
}
