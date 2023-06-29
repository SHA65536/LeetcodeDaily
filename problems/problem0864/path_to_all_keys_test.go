package problem0864

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Grid     []string
	Expected int
}

var TestCases = []TestCase{
	{
		[]string{
			"@.a..",
			"###.#",
			"b.A.B",
		}, 8,
	},
	{
		[]string{
			"@..aA",
			"..B#.",
			"....b",
		}, 6,
	},
	{
		[]string{
			"@Aa",
		}, -1,
	},
}

func TestShortestPathToAllKeys(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := shortestPathAllKeys(tc.Grid)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkShortestPathToAllKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			shortestPathAllKeys(tc.Grid)
		}
	}
}
