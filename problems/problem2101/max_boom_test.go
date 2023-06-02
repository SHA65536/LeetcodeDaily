package problem2101

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	Expected int
}

var TestCases = []TestCase{
	{[][]int{{2, 1, 3}, {6, 1, 4}}, 2},
	{[][]int{{1, 1, 5}, {10, 10, 5}}, 1},
	{[][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}, 5},
	{[][]int{{855, 82, 158}, {17, 719, 430}, {90, 756, 164}, {376, 17, 340}, {691, 636, 152}, {565, 776, 5}, {464, 154, 271}, {53, 361, 162}, {278, 609, 82}, {202, 927, 219}, {542, 865, 377}, {330, 402, 270}, {720, 199, 10}, {986, 697, 443}, {471, 296, 69}, {393, 81, 404}, {127, 405, 177}}, 9},
}

func TestMaxBoom(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maximumDetonation(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxBoom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maximumDetonation(tc.Input)
		}
	}
}
