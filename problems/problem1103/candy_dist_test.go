package problem1103

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Candies  int
	People   int
	Expected []int
}

var Results = []Result{
	{7, 4, []int{1, 2, 3, 1}},
	{10, 3, []int{5, 2, 3}},
}

func TestDistributeCandies(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := distributeCandies(res.Candies, res.People)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
