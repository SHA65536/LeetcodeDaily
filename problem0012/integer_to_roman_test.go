package problem0012

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected string
}

var Results = []Result{
	{3, "III"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
	{37, "XXXVII"},
	{99, "XCIX"},
	{333, "CCCXXXIII"},
	{999, "CMXCIX"},
}

func TestIntToRoman(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, intToRoman(res.Input), fmt.Sprintf("%+v", res))
	}
}
