package problem1361

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N           int
	Left, Right []int
	Expected    bool
}

var TestCases = []TestCase{
	//{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
	//{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
	//{2, []int{1, 0}, []int{-1, -1}, false},
	{4, []int{1, 2, 0, -1}, []int{-1, -1, -1, -1}, false},
}

func TestValidateBinaryTree(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := validateBinaryTreeNodes(tc.N, tc.Left, tc.Right)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkValidateBinaryTree(b *testing.B) {
	for _, tc := range TestCases {
		for i := 0; i < b.N; i++ {
			validateBinaryTreeNodes(tc.N, tc.Left, tc.Right)
		}
	}
}
