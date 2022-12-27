package problem2279

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cap        []int
	Rock       []int
	Additional int
	Expected   int
}

var Results = []Result{
	{[]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2, 3},
	{[]int{10, 2, 2}, []int{2, 2, 0}, 100, 3},
}

func TestMaximumFilledBags(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximumBags(res.Cap, res.Rock, res.Additional)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
