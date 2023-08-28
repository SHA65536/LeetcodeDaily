package problem0225

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Cmds     []string
	Values   []int
	Expected []int
}

var TestCases = []TestCase{
	{
		Cmds:     []string{"MyStack", "push", "push", "top", "pop", "empty"},
		Values:   []int{0, 1, 2, 0, 0, 0},
		Expected: []int{0, 0, 0, 2, 2, 0},
	},
}

func TestStackFromQueues(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var stack MyStack
		want := tc.Expected
		got := make([]int, len(want))
		for i, cmd := range tc.Cmds {
			switch cmd {
			case "MyStack":
				stack = Constructor()
			case "push":
				stack.Push(tc.Values[i])
			case "pop":
				got[i] = stack.Pop()
			case "top":
				got[i] = stack.Top()
			case "empty":
				if stack.Empty() {
					got[i] = -1
				}
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLongestPairChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			var stack MyStack
			for i, cmd := range tc.Cmds {
				switch cmd {
				case "MyStack":
					stack = Constructor()
				case "push":
					stack.Push(tc.Values[i])
				case "pop":
					stack.Pop()
				case "top":
					stack.Top()
				case "empty":
					stack.Empty()
				}
			}
		}
	}
}
