package problem0557

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	{"God Ding", "doG gniD"},
}

func TestReverseWords(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseWords(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
