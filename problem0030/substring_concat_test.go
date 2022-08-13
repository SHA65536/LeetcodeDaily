package problem0030

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Words    []string
	Expected []int
}

var Results = []Result{
	{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
	{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
	{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
}

func TestFindSubstringConcat(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findSubstring(res.Input, res.Words)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
