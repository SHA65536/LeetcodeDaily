package problem0881

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Limit    int
	Expected int
}

var Results = []Result{
	{[]int{1, 2}, 3, 1},
	{[]int{3, 2, 2, 1}, 3, 3},
	{[]int{3, 5, 3, 4}, 5, 4},
}

func TestLifeBoats(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		temp := make([]int, len(res.Input))
		copy(temp, res.Input)
		got := numRescueBoats(temp, res.Limit)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
