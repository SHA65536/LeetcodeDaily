package problem0743

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Times    [][]int
	N        int
	K        int
	Expected int
}

var Results = []Result{
	{
		Times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
		N:     4, K: 2, Expected: 2,
	},
	{
		Times: [][]int{{1, 2, 1}},
		N:     2, K: 1, Expected: 1,
	},
	{
		Times: [][]int{{1, 2, 1}},
		N:     2, K: 2, Expected: -1,
	},
	{
		Times: [][]int{{1, 2, 1}, {2, 1, 3}},
		N:     2, K: 2, Expected: 3,
	},
	{
		Times: [][]int{{1, 1, 1}},
		N:     1, K: 1, Expected: 0,
	},
	{
		Times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}},
		N:     3, K: 1, Expected: 2,
	},
}

func TestNetworkDelayTime(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := networkDelayTime(res.Times, res.N, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

/*
Constraints:
1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= ui, vi <= n
ui != vi
0 <= wi <= 100
All the pairs (ui, vi) are unique. (i.e., no multiple edges.)
*/
