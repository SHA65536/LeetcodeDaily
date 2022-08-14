package problem0126

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Begin, End string
	Words      []string
	Expected   [][]string
}

var Results = []Result{
	{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"},
		[][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}},
	{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"},
		[][]string{}},
}

func TestWordLadderII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findLadders(res.Begin, res.End, res.Words)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
