package problem0376

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{1, 7, 4, 9, 2, 5}, 6},
	{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	{[]int{912, 158, 823, 462, 291, 702, 965, 450, 221, 501, 13, 722, 780, 332, 166, 653, 294, 458, 547, 802}, 13},
}

func TestWiggleMaxLength(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := wiggleMaxLength(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
