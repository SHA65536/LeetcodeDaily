package problem0491

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected [][]int
}

var Results = []Result{
	{[]int{4, 6, 7, 7}, [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}}},
	{[]int{4, 4, 3, 2, 1}, [][]int{{4, 4}}},
}

func TestNonDecreasingSubsequences(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findSubsequences(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
