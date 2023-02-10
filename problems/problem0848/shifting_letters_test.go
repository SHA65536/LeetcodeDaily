package problem0848

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Shifts   []int
	Expected string
}

var Results = []Result{
	{"abc", []int{3, 5, 9}, "rpl"},
	{"aaa", []int{1, 2, 3}, "gfd"},
	{"abcdefghijklmnopqrstuzwxyz", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, "nnmkhdysldukznamxhqyfpquxz"},
}

func TestShiftingLetters(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := shiftingLetters(res.Input, res.Shifts)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
