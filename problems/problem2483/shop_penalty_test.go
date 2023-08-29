package problem2483

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Expected int
}

var TestCases = []TestCase{
	{"YYNY", 2},
	{"NNNNN", 0},
	{"YYYY", 4},
}

func TestMinShopPenalty(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := bestClosingTime(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinShopPenalty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			bestClosingTime(tc.Input)
		}
	}
}
