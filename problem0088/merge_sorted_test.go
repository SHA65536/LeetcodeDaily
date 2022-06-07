package problem0088

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums1, Nums2 []int
	M, N         int
	Expected     []int
}

var Results = []Result{
	{
		[]int{1, 2, 3, 0, 0, 0},
		[]int{2, 5, 6},
		3, 3,
		[]int{1, 2, 2, 3, 5, 6},
	},
	{
		[]int{1, 2, 3, 0, 0, 0, 0},
		[]int{2, 5, 6, 7},
		3, 4,
		[]int{1, 2, 2, 3, 5, 6, 7},
	},
	{
		[]int{1, 2, 3, 4, 5, 0, 0},
		[]int{2, 5},
		5, 2,
		[]int{1, 2, 2, 3, 4, 5, 5},
	},
	{
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{1, 2, 2, 3, 5, 6, 7},
		0, 7,
		[]int{1, 2, 2, 3, 5, 6, 7},
	},
	{
		[]int{1, 2, 3, 0, 0, 0},
		[]int{4, 5, 6},
		3, 3,
		[]int{1, 2, 3, 4, 5, 6},
	},
}

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		merge(res.Nums1, res.M, res.Nums2, res.N)
		assert.Equal(want, res.Nums1, fmt.Sprintf("%+v", res))
	}
}
