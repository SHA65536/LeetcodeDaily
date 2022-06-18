package problem0745

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Words          []string
	Prefix, Suffix string
	Expected       int
}

var Results = []Result{
	{
		Words:  []string{"apple"},
		Prefix: "a", Suffix: "e",
		Expected: 0,
	},
	{
		Words:  []string{"apple", "able"},
		Prefix: "a", Suffix: "e",
		Expected: 1,
	},
	{
		Words:  []string{"adrian", "apple", "able"},
		Prefix: "a", Suffix: "e",
		Expected: 2,
	},
	{
		Words:  []string{"a", "a", "a", "b"},
		Prefix: "a", Suffix: "a",
		Expected: 2,
	},
}

func TestWordFilter(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		wf := Constructor(res.Words)
		want := res.Expected
		got := wf.F(res.Prefix, res.Suffix)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
