package problem0048

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected [][]int
}

var Results = []Result{
	{[][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	},
		[][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
	{[][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	},
		[][]int{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		}},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	},
		[][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
	{[][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	},
		[][]int{
			{21, 16, 11, 6, 1},
			{22, 17, 12, 7, 2},
			{23, 18, 13, 8, 3},
			{24, 19, 14, 9, 4},
			{25, 20, 15, 10, 5},
		}},
}

func TestRotateImage(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		rotate(res.Input)
		assert.Equal(want, res.Input, fmt.Sprintf("%+v", res))
	}
}
