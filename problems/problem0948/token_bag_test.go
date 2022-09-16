package problem0948

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Tokens   []int
	Power    int
	Expected int
}

var Results = []Result{
	{[]int{100}, 50, 0},
	{[]int{100, 200}, 150, 1},
	{[]int{100, 200, 300, 400}, 200, 2},
	{[]int{31, 46, 30, 94, 40, 26, 10, 64, 66, 23, 54, 56, 60, 27, 9, 59, 92, 4, 84, 49, 24, 91, 93, 21, 46, 8, 4, 30, 28, 82, 68, 48, 19, 95, 15, 60, 93, 48, 92, 50, 98, 48, 81, 8, 39, 20, 42, 58, 84, 21, 90, 36, 46, 90, 23, 96, 86, 52, 5, 79, 22, 25, 68, 86, 87, 44, 84, 30, 71, 43, 18, 31, 83, 71, 53, 76, 60, 68, 27, 90, 51, 48, 30, 7, 24, 59, 1, 70, 47, 13, 68, 41, 79, 58, 60, 10, 22, 17, 56, 91}, 100, 42},
}

func TestBagOfTokens(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := bagOfTokensScore(res.Tokens, res.Power)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
