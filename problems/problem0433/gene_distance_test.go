package problem0433

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	End      string
	Bank     []string
	Expected int
}

var Results = []Result{
	{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
	{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
	{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
	{"AACCGGTT", "AACCGGTA", []string{}, -1},
}

func TestMinGeneMutation(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minMutation(res.Input, res.End, res.Bank)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
