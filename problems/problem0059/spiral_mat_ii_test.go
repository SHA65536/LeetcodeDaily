package problem0059

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected [][]int
}

var Results = []Result{
	{1, [][]int{{1}}},
	{2, [][]int{
		{1, 2},
		{4, 3}}},
	{3, [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5}}},
	{4, [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7}}},
}

func TestSpiralMatrixII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := generateMatrix(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
