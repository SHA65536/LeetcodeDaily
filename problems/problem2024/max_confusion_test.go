package problem2024

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Answer   string
	K        int
	Expected int
}

var TestCases = []TestCase{
	{"TTFF", 2, 4},
	{"TFFT", 1, 3},
	{"TTFTTFTT", 2, 8},
	{"TTFTTFTT", 1, 5},
}

func TestMaxConfusion(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxConsecutiveAnswers(tc.Answer, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxConfusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxConsecutiveAnswers(tc.Answer, tc.K)
		}
	}
}
