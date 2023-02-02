package problem0953

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Words    []string
	Order    string
	Expected bool
}

var Results = []Result{
	{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
	{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
	{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	{[]string{"ubg", "kwh"}, "qcipyamwvdjtesbghlorufnkzx", true},
}

func TestVerifyDictionary(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := isAlienSorted(res.Words, res.Order)
		want := res.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
