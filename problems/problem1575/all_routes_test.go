package problem1575

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Cities              []int
	Start, Finish, Fuel int
	Expected            int
}

var TestCases = []TestCase{
	{[]int{2, 3, 6, 8, 4}, 1, 3, 5, 4},
	{[]int{4, 3, 1}, 1, 0, 6, 5},
	{[]int{5, 2, 1}, 0, 2, 3, 0},
}

func TestTallestBillboard(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := countRoutes(tc.Cities, tc.Start, tc.Finish, tc.Fuel)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTallestBillboard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			countRoutes(tc.Cities, tc.Start, tc.Finish, tc.Fuel)
		}
	}
}
