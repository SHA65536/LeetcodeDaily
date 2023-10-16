package problem0119

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    int
	Expected []int
}

var TestCases = []TestCase{
	{3, []int{1, 3, 3, 1}},
	{0, []int{1}},
	{1, []int{1, 1}},
	{10, []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1}},
	{33, []int{1, 33, 528, 5456, 40920, 237336, 1107568, 4272048, 13884156, 38567100, 92561040, 193536720, 354817320, 573166440, 818809200, 1037158320, 1166803110, 1166803110, 1037158320, 818809200, 573166440, 354817320, 193536720, 92561040, 38567100, 13884156, 4272048, 1107568, 237336, 40920, 5456, 528, 33, 1}},
}

func TestGetPascalRowII(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := getRow(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkGetPascalRowII(b *testing.B) {
	for _, tc := range TestCases {
		for i := 0; i < b.N; i++ {
			getRow(tc.Input)
		}
	}
}
