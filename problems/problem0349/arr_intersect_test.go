package problem0349

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums1, Nums2 []int
	Expected     []int
}

var Results = []Result{
	{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
	{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
}

func TestArrayIntersection(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := intersection(res.Nums1, res.Nums2)
		sort.Ints(got)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
