package problem0863

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    *TreeNode
	Target   *TreeNode
	K        int
	Expected []int
}

var TestCases = []TestCase{
	{
		MakeTree(3, 5, 6, NULL, NULL, 2, 7, NULL, NULL, 4, NULL, NULL, 1, 0, NULL, NULL, 8, NULL, NULL),
		nil, 2, []int{7, 4, 1},
	},
	{
		MakeTree(1, NULL, NULL),
		nil, 3, []int{},
	},
}

func init() {
	TestCases[0].Target = TestCases[0].Input.Left
	TestCases[1].Target = TestCases[1].Input
}

func TestDistanceKTree(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := distanceK(tc.Input, tc.Target, tc.K)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkDistanceKTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			distanceK(tc.Input, tc.Target, tc.K)
		}
	}
}
