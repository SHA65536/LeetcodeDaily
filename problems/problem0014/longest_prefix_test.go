package problem0014

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected string
}

var Results = []Result{
	{[]string{"flow", "flower", "fly"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"fly", "flower"}, "fl"},
	{[]string{"floop", "floopy", "floopendium"}, "floop"},
	{[]string{"", "abc"}, ""},
}

func TestLongestPrefix(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, longestCommonPrefix(res.Input), fmt.Sprintf("%+v", res))
	}
}

/*
Constraints:
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.
*/
