package problem0974

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected int
}

var Results = []Result{
	{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
	{[]int{5}, 9, 0},
}

func TestSubbarryDivisibleK(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := subarraysDivByK(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
