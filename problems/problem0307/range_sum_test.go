package problem0307

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Query    []int
	Update   []int
	Expected int
}

var Results = []Result{
	{[]int{0, 2}, nil, 9},
	{nil, []int{1, 2}, 0},
	{[]int{0, 2}, nil, 8},
}

func TestSumRange(t *testing.T) {
	assert := assert.New(t)
	var matrix = []int{1, 3, 5}

	mat := Constructor(matrix)
	for _, res := range Results {
		if res.Query != nil {
			want := res.Expected
			got := (&mat).SumRange(res.Query[0], res.Query[1])
			assert.Equal(want, got, fmt.Sprintf("%+v", res))
		} else {
			mat.Update(res.Update[0], res.Update[1])
		}
	}
}
