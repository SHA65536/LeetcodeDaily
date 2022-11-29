package problem0424

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	K        int
	Expected int
}

var Results = []Result{
	{"ABAB", 2, 4},
	{"AABABBA", 1, 4},
}

func TestLongestRepeatingCharacterRep(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := characterReplacement(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
