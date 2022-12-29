package problem1834

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected []int
}

var Results = []Result{
	/*{[][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}, []int{0, 2, 3, 1}},
	{[][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}, []int{4, 3, 2, 0, 1}},*/
	{[][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}, []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
}

func TestSingleThreadCPUOrder(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := getOrder(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
