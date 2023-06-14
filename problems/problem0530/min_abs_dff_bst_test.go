package problem0530

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    *TreeNode
	Expected int
}

var TestCases = []TestCase{
	{MakeTree(4, 2, 1, NULL, NULL, 3, NULL, NULL, 6, NULL, NULL), 1},
	{MakeTree(1, 0, NULL, NULL, 48, 12, NULL, NULL, 49, NULL, NULL), 1},
}

func TestMinDiffInBSTPointers(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := getMinimumDifferencePointers(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinDiffInBSTPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			getMinimumDifferencePointers(tc.Input)
		}
	}
}

func TestMinDiffInBSTChans(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := getMinimumDifferenceChan(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinDiffInBSTChans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			getMinimumDifferenceChan(tc.Input)
		}
	}
}
