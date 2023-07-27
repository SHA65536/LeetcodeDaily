package problem2141

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N         int
	Batteries []int
	Expected  int64
}

var TestCases = []TestCase{
	{2, []int{3, 3, 3}, 4},
	{2, []int{1, 1, 1, 1}, 2},
}

func TestMaxBatteryLife(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxRunTime(tc.N, tc.Batteries)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxBatteryLife(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxRunTime(tc.N, tc.Batteries)
		}
	}
}
func TestMaxBatteryLifeBinSearch(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxRunTimeBinSearch(tc.N, tc.Batteries)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxBatteryLifeBinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxRunTimeBinSearch(tc.N, tc.Batteries)
		}
	}
}
