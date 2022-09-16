package problem0013

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
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"XXXVII", 37},
	{"XCIX", 99},
	{"CCCXXXIII", 333},
	{"CMXCIX", 999},
}

func TestRomanToInt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, romanToInt(res.Input), fmt.Sprintf("%+v", res))
	}
}
