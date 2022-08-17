package problem0804

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
	{[]string{"gin", "zen", "gig", "msg"}, 2},
	{[]string{"a"}, 1},
}

func TestUniqueMorse(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := uniqueMorseRepresentations(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
