package problem0004

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums1    []int
	Nums2    []int
	Expected float64
}

var Results = []Result{
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1}, []int{}, 1},
	{[]int{}, []int{2}, 2},
	{[]int{1, 3, 6, 7}, []int{}, 4.5},
	{[]int{1, 3, 6, 7, 78, 687, 1111}, []int{4, 89, 784, 100000, 2569874}, 83.5},
	{[]int{1, 3, 6, 7, 9, 15, 20, 23, 35}, []int{8957648}, 12},
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findMedianSortedArrays(res.Nums1, res.Nums2)
		assert.Equal(want, got, fmt.Sprint(res.Nums1, res.Nums2))
	}
}
