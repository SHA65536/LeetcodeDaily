package problem1980

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input []string
}

var TestCases = []TestCase{
	{[]string{"01", "10"}},
	{[]string{"00", "10"}},
	{[]string{"111", "011", "001"}},
}

func TestUniqueBinaryNumber(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var got = findDifferentBinaryString(tc.Input)
		assert.Equal(len(tc.Input), len(got), fmt.Sprintf("%+v", tc))
		assert.NotContains(tc.Input, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkUniqueBinaryNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findDifferentBinaryString(tc.Input)
		}
	}
}
