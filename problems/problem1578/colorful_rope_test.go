package problem1578

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Colors   string
	Times    []int
	Expected int
}

var Results = []Result{
	{"abaac", []int{1, 2, 3, 4, 5}, 3},
	{"abc", []int{1, 2, 3}, 0},
	{"aabaa", []int{1, 2, 3, 4, 1}, 2},
	{"aaabc", []int{2, 1, 5, 7, 8}, 3},
	{"bcaaa", []int{7, 8, 5, 1, 2}, 3},
}

func TestColorfulRope(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minCost(res.Colors, res.Times)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
