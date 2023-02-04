package problem2553

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected []int
}

var Results = []Result{
	{[]int{13, 25, 83, 77}, []int{1, 3, 2, 5, 8, 3, 7, 7}},
	{[]int{7, 1, 3, 9}, []int{7, 1, 3, 9}},
	{[]int{12345}, []int{1, 2, 3, 4, 5}},
	{[]int{1234}, []int{1, 2, 3, 4}},
}

func TestMostBeautifulItem(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := separateDigits(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
