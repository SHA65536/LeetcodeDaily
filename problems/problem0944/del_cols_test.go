package problem0944

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int
}

var Results = []Result{
	{[]string{"abc", "bce", "cae"}, 1},
	{[]string{"cba", "daf", "ghi"}, 1},
	{[]string{"a", "b"}, 0},
	{[]string{"zyx", "wvu", "tsr"}, 3},
}

func TestDeleteCollumnsNotSorted(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDeletionSize(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
