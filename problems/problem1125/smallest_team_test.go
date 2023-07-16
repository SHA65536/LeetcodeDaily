package problem1125

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Req      []string
	People   [][]string
	Expected []int
}

var TestCases = []TestCase{
	{
		[]string{"java", "nodejs", "reactjs"},
		[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
		[]int{0, 2},
	},
	{
		[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		[][]string{
			{"algorithms", "math", "java"},
			{"algorithms", "math", "reactjs"},
			{"java", "csharp", "aws"},
			{"reactjs", "csharp"},
			{"csharp", "math"},
			{"aws", "java"},
		},
		[]int{1, 2},
	},
}

func TestSmallestTeam(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := smallestSufficientTeam(tc.Req, tc.People)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSmallestTeam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			smallestSufficientTeam(tc.Req, tc.People)
		}
	}
}
