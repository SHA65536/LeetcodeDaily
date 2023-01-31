package problem1626

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Score    []int
	Age      []int
	Expected int
}

var Results = []Result{
	{[]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}, 34},
	{[]int{4, 5, 6, 5}, []int{2, 1, 2, 1}, 16},
	{[]int{1, 2, 3, 5}, []int{8, 9, 10, 1}, 6},
	{[]int{319776, 611683, 835240, 602298, 430007, 574, 142444, 858606, 734364, 896074}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 5431066},
}

func TestBestBasketballTeam(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := bestTeamScore(res.Score, res.Age)
		want := res.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
