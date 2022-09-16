package problem0967

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N, K     int
	Expected []int
}

var Results = []Result{
	{3, 7, []int{181, 292, 707, 818, 929}},
	{2, 1, []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
	{2, 0, []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
}

func TestAverageOfLevels(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numsSameConsecDiff(res.N, res.K)
		sort.Ints(got)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
