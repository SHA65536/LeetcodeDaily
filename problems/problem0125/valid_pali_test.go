package problem0125

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected bool
}

var Results = []Result{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
}

func TestValidPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPalindrome(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
