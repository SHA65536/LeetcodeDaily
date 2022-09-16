package problem0022

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected []string
}

var Results = []Result{
	{1, []string{"()"}},
	{2, []string{"(())", "()()"}},
	{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
}

func TestMergeTwoLists(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := generateParenthesis(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
