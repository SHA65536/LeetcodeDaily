package problem0508

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    *TreeNode
	Expected []int
}

var TestCases = []TestCase{
	{MakeTree(5, 2, NULL, NULL, -3, NULL, NULL), []int{2, -3, 4}},
	{MakeTree(5, 2, NULL, NULL, -5, NULL, NULL), []int{2}},
}

func TestTopFreqSubtreeSum(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := findFrequentTreeSum(tc.Input)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTopFreqSubtreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findFrequentTreeSum(tc.Input)
		}
	}
}
