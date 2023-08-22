package problem0168

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    int
	Expected string
}

var TestCases = []TestCase{
	{1, "A"},
	{2, "B"},
	{26, "Z"},
	{27, "AA"},
	{28, "AB"},
	{701, "ZY"},
	{15489, "VWS"},
	{2147483647, "FXSHRXW"},
}

func TestConvertToExcel(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := convertToTitle(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkConvertToExcel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			convertToTitle(tc.Input)
		}
	}
}
