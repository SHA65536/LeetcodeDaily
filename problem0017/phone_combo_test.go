package problem0017

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected []string
}

var Results = []Result{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"", []string{}},
	{"2", []string{"a", "b", "c"}},
}

func TestLetterCombinations(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, letterCombinations(res.Input), fmt.Sprintf("%+v", res))
	}
}

/*
Constraints:
0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/
