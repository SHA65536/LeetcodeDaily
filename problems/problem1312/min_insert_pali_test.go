package problem1312

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"zzazz", 0},
	{"mbadm", 2},
	{"leetcode", 5},
}

func TestNumberOfInsertionsForPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minInsertions(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
